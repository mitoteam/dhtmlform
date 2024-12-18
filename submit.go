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

			if control.data.value != nil {
				tag.Attribute("value", mttools.AnyToString(control.data.value))
			}

			if control.label.IsEmpty() {
				tag.Append("Submit")
			} else {
				tag.Append(control.label)
			}

			out.Append(tag)
			return out
		},
	})
}

func NewSubmitBtn() *FormControlElement {
	return NewFormControl(submitControlKind, "submit")
}
