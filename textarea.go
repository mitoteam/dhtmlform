package dhtmlform

import (
	"github.com/mitoteam/dhtml"
)

const textareaControlKind = "textarea"

func init() {
	RegisterFormControlHandler(textareaControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div()

			if !control.label.IsEmpty() {
				rootTag.Append(dhtml.NewTag("label").Attribute("for", control.GetId()).
					Class("fc-label").Append(control.label))
			}

			rootTag.Append(dhtml.NewTag("textarea").Id(control.GetId()).Append(control.data.value))

			if !control.note.IsEmpty() {
				rootTag.Append(dhtml.Div().Append(dhtml.NewTag("small").Append(control.note)))
			}

			out.Append(rootTag)
			return out
		},
	})
}

func NewTextarea(name string) *FormControlElement {
	return NewFormControl(textareaControlKind, name)
}
