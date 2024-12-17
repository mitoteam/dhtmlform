package dhtmlform

import (
	"github.com/mitoteam/dhtml"
)

const textareaControlKind = "textarea"

func init() {
	RegisterFormControlHandler(textareaControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			tag := dhtml.NewTag("textarea").Append(control.GetValue())

			out.Append(tag)

			if !control.note.IsEmpty() {
				out.Append(dhtml.Div().Class("fc-note").Append(control.note))
			}

			return out
		},
	})
}

func NewTextarea(name string) *FormControlElement {
	return NewFormControl(textareaControlKind, name)
}
