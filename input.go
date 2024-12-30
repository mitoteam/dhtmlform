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

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Attribute("type", control.GetProp("type").(string)).
				Attribute("name", control.GetName()).Attribute("value", mttools.AnyToString(control.GetValue()))

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
	return NewFormControl(inputControlKind, name).SetProp("type", "text")
}

func NewPasswordInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "password")
}

func NewEmailInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "email")
}

func NewDateInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "date")
}

func NewNumberInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "number")
}

func NewTelInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "tel")
}

func NewTimeInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "time")
}

func NewUrlInput(name string) *FormControlElement {
	return NewFormControl(inputControlKind, name).SetProp("type", "url")
}
