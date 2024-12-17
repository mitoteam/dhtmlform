package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const textareaControlKind = "textarea"

func init() {
	RegisterFormControlHandler(textareaControlKind, &FormControlHandler{
		RenderF: func(props mttools.Values) (out dhtml.HtmlPiece) {
			out.Append("TEXTAREA!")

			return out
		},
	})
}

func NewTextarea(name string) *FormControlElement {
	return NewFormControl(textareaControlKind, name)
}
