package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

type FormControlElement struct {
	controlKind string
	name        string

	//additional properties
	props mttools.Values
}

var _ dhtml.ElementI = (*FormControlElement)(nil)

func NewFormControl(controlKind string, name string) *FormControlElement {
	if _, ok := GetFormControlHandler(controlKind); !ok {
		return nil
	}

	return &FormControlElement{
		controlKind: controlKind,
		name:        name,
		props:       mttools.NewValues(),
	}
}

func (e *FormControlElement) GetTags() dhtml.TagList {
	var out dhtml.HtmlPiece

	if handler, ok := GetFormControlHandler(e.controlKind); ok {
		out.Append(handler.RenderF(e.props))
	}

	return out.GetTags()
}
