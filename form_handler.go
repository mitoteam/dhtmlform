package dhtmlform

import (
	"net/http"

	"github.com/mitoteam/dhtml"
)

type FormHandler struct {
	RenderF   func(formBody *dhtml.HtmlPiece, fd *FormData)
	ValidateF func(fd *FormData)
	SubmitF   func(fd *FormData)
}

const hiddenBuildIdFieldName = "dhtmlform_build_id"

func (fh *FormHandler) Render(fc *FormContext) *dhtml.HtmlPiece {
	var formBody dhtml.HtmlPiece

	var fd = formDataFromPOST(fc)

	if fd != nil { //found in POST values and data store
		fd.redirectUrl = ""
		fd.rebuild = false
		//TODO: clear errors

		fh.ValidateF(fd)

		if true { //TODO: fd.HasError() {
			//TODO: formOut.Append(settings.FormErrorsRendererF(fd))
			fd.rebuild = true //and display form again
		} else {
			//there were no errors
			fh.SubmitF(fd)
		}

		if !fd.rebuild {
			delete(formDataStore, fd.build_id)

			//check redirect (first from FormData, then from FormContext)
			var redirectUrl = fd.redirectUrl

			if redirectUrl == "" {
				redirectUrl = fc.redirectUrl
			}

			if redirectUrl != "" {
				http.Redirect(fc.w, fc.r, redirectUrl, http.StatusSeeOther)
				return dhtml.NewHtmlPiece() //empty html
			}

			//we are not rebuilding, not redirecting = new form should be built from scratch
			fd = nil
		}
	}

	if fd == nil {
		fd = NewFormData()
		fd.args.CopyFrom(fc.args)
		fd.params.CopyFrom(fc.params)
	}

	//<form> tag
	form := dhtml.NewForm().
		Append(NewHidden(hiddenBuildIdFieldName).Default(fd.build_id))

	if fh.RenderF != nil {
		fh.RenderF(&formBody, fd)
		formBody.WalkR(fd.processControlDataWalkerF)
	}

	//save to store for rebuilds
	formDataStore[fd.build_id] = fd

	form.Append(formBody)

	//wrap it all into container <div>
	div := dhtml.Div().Class("dhtml-form").Append(form)
	return dhtml.Piece(div)
}

// Check build_id in POST values, try to find form data in data store, re-hydrate it with submitted values and return.
// Returns nil if there is no FormData to rebuild.
func formDataFromPOST(fc *FormContext) (fd *FormData) {
	build_id := fc.r.PostFormValue("form_build_id")

	// check if it is being re-build (from POST request)
	if build_id == "" {
		return nil
	}

	// check if build_id is in store
	fd, ok := formDataStore[build_id]
	if !ok {
		return nil
	}

	//refresh params from context
	fd.params.CopyFrom(fc.params)

	//re-hydrate form_data.values from POST data
	for name, controlData := range fd.controlsData {
		controlData.value = nil

		if rawPostValue, ok := fc.r.PostForm[name]; ok {
			if len(rawPostValue) == 1 { //array of single element, just take first one
				controlData.value = rawPostValue[0]
			} else {
				controlData.value = rawPostValue
			}
		}

		if handler, ok := GetFormControlHandler(controlData.controlKind); ok {
			if handler.ProcessPostValueF != nil {
				controlData.value = handler.ProcessPostValueF(controlData.value)
			}
		}

		//set it back to FormData
		fd.controlsData[name] = controlData
	}

	return fd
}

// Walker function to set control values from FormData if form being rebuild after post
func (fd *FormData) processControlDataWalkerF(e dhtml.ElementI, args ...any) {
	if control, ok := e.(*FormControlElement); ok {
		if storedData, ok := fd.controlsData[control.name]; ok && storedData.controlKind == control.data.controlKind {
			control.data = storedData
		} else {
			fd.controlsData[control.name] = control.data
		}

		//control.Note(fmt.Sprintf("DBG: processControlDataWalkerF Walked, rebuild: %t", fd.rebuild))
	}
}
