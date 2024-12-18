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

	var fd = fc.getFormDataFromPOST()

	if fd != nil { //found in POST values and data store
		fd.redirectUrl = ""
		fd.rebuild = false

		//basic internal validations (like required values)
		fd.ClearErrors()
		fd.validateFormControls()

		// custom form handler validations
		if fh.ValidateF != nil {
			fh.ValidateF(fd)
		}

		// render errors if any
		if fd.HasError() {
			rootTag.Append(settings.FormErrorsRenderF(&fd.errors))
			fd.rebuild = true //and display form again
		}

		// no rebuild requested, do submit
		if !fd.rebuild {
			if fh.SubmitF != nil {
				fh.SubmitF(fd)
			}
		}

		//rebuilt flag can be set in SubmitF() so check it again
		if !fd.rebuild {
			delete(formDataStore, fd.build_id)

			//check redirect (first from FormData, then from FormContext)
			var redirectUrl = fd.redirectUrl

			if redirectUrl == "" {
				redirectUrl = fc.redirectUrl
			}

			if redirectUrl != "" {
				http.Redirect(fc.w, fc.r, redirectUrl, http.StatusSeeOther)
				return dhtml.NewHtmlPiece() //empty piece
			}

			//we are not rebuilding and not redirecting = new form should be built from scratch
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

	rootTag.Append(form)
	return dhtml.Piece(rootTag)
}
