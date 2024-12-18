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
	rootTag := dhtml.Div().Class("dhtml-form")
	var formBody dhtml.HtmlPiece

	var fd = fc.formDataFromPOST()

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

	form.Append(dhtml.Dbg("%+v", fd.controlsData))

	if fh.RenderF != nil {
		fh.RenderF(&formBody, fd)
		formBody.WalkR(fd.processControlDataWalkerF)
	}

	//save to store for rebuilds
	formDataStore[fd.build_id] = fd

	form.Append(formBody)

	rootTag.Append(form)
	return dhtml.Piece(rootTag)
}
