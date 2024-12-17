package dhtmlform

import "github.com/mitoteam/dhtml"

type FormHandler struct {
	RenderF   func(formBody *dhtml.HtmlPiece, fd *FormData)
	ValidateF func(fd *FormData)
	SubmitF   func(fd *FormData)
}

func (f *FormHandler) Render(fc *FormContext) *dhtml.HtmlPiece {
	var formBody dhtml.HtmlPiece

	var fd *FormData
	fd = &FormData{}
	if f.RenderF != nil {
		f.RenderF(&formBody, fd)
	}

	//wrap it all into container <div>
	div := dhtml.Div().Class("dhtml-form").Append(formBody)

	return dhtml.Piece(div)
}
