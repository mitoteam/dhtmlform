package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const submitControlKind = "submit_btn"

func init() {
	RegisterFormControlHandler(submitControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			tag := dhtml.NewTag("button").Attribute("type", "submit")

			if !mttools.IsEmpty(control.GetValue()) {
				tag.Attribute("name", control.GetName()).Attribute("value", mttools.AnyToString(control.GetValue()))
			}

			if control.GetLabel().IsEmpty() {
				tag.Append("Submit")
			} else {
				tag.Append(control.GetLabel())
			}

			out.Append(tag)
			return out
		},
	})
}

func NewSubmitBtn() *FormControlElement {
	return NewFormControl(submitControlKind, "submit")
}
