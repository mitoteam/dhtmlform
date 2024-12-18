package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const checkboxControlKind = "checkbox"
const checkboxCheckedFormValue = "on"

func init() {
	RegisterFormControlHandler(checkboxControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div()

			if control.IsError() {
				rootTag.Styles(errorBlockStyle)
			}

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Attribute("type", "checkbox").
				Attribute("name", control.Name).Attribute("value", checkboxCheckedFormValue)

			if mttools.AnyToBool(control.GetValue()) {
				inputTag.Attribute("checked", "")
			}

			rootTag.Append(inputTag)

			if !control.GetLabel().IsEmpty() {
				labelOut := dhtml.NewLabel().For(control.GetId()).Append(control.GetLabel())

				if control.IsRequired() {
					labelOut.Append(dhtml.Span().Styles("color: red; font-weight: bolder;").Text("*"))
				}

				rootTag.Append(labelOut)
			}

			if !control.note.IsEmpty() {
				rootTag.Append(control.renderNote())
			}

			out.Append(rootTag)
			return out
		},

		ProcessPostValueF: func(rawValue any) any {
			if rawValue == checkboxCheckedFormValue {
				return true
			} else {
				return mttools.AnyToBool(rawValue)
			}
		},
	})
}

func NewCheckbox(name string) *FormControlElement {
	return NewFormControl(checkboxControlKind, name)
}
