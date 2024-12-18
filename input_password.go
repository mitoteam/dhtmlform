package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const passwordControlKind = "passwordinput"

func init() {
	RegisterFormControlHandler(passwordControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div()

			if control.IsError() {
				rootTag.Styles(errorBlockStyle)
			}

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(control.renderLabel())
			}

			rootTag.Append(
				dhtml.NewTag("input").Id(control.GetId()).Attribute("type", "password").
					Attribute("name", control.Name).Attribute("value", mttools.AnyToString(control.GetValue())),
			)

			if !control.note.IsEmpty() {
				rootTag.Append(control.renderNote())
			}

			out.Append(rootTag)
			return out
		},
	})
}

func NewPasswordInput(name string) *FormControlElement {
	return NewFormControl(passwordControlKind, name)
}
