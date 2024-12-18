package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const inputControlKind = "textinput"

func init() {
	RegisterFormControlHandler(inputControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div()

			if control.IsError() {
				rootTag.Styles(errorBlockStyle)
			}

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(control.renderLabel())
			}

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Attribute("type", "text").
				Attribute("name", control.Name).Attribute("value", mttools.AnyToString(control.GetValue()))

			if control.GetPlaceholder() != "" {
				inputTag.Attribute("placeholder", control.GetPlaceholder())
			}

			rootTag.Append(inputTag)

			if !control.note.IsEmpty() {
				rootTag.Append(control.renderNote())
			}

			out.Append(rootTag)
			return out
		},
	})
}

func NewTextInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name)
}
