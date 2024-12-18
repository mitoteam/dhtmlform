package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

const hiddenControlKind = "hidden"

func init() {
	RegisterFormControlHandler(hiddenControlKind, &FormControlHandler{
		RenderF: func(control *FormControlElement) (out dhtml.HtmlPiece) {
			out.Append(
				dhtml.NewTag("input").Attribute("type", "hidden").
					Attribute("name", control.name).
					Attribute("value", mttools.AnyToString(control.data.value)),
			)

			return out
		},
	})
}

func NewHidden(name string) *FormControlElement {
	return NewFormControl(hiddenControlKind, name)
}
