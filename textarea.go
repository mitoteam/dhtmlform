package dhtmlform

import (
	"github.com/mitoteam/dhtml"
)

const textareaControlKind = "textarea"

func init() {
	RegisterFormControlHandler(textareaControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div()

			if control.IsError() {
				rootTag.Styles(errorBlockStyle)
			}

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(control.renderLabel())
			}

			textareaTag := dhtml.NewTag("textarea").Id(control.GetId()).Attribute("name", control.Name).
				Append(control.GetValue())

			if control.GetPlaceholder() != "" {
				textareaTag.Attribute("placeholder", control.GetPlaceholder())
			}

			rootTag.Append(textareaTag)

			if !control.note.IsEmpty() {
				rootTag.Append(control.renderNote())
			}

			out.Append(rootTag)
			return out
		},
	})
}

func NewTextarea(name string) *FormControlElement {
	return NewFormControl(textareaControlKind, name)
}
