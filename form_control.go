package dhtmlform

import (
	"log"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

type FormControlElement struct {
	controlKind string
	props       mttools.Values
}

var _ dhtml.ElementI = (*FormControlElement)(nil)

func NewFormControl(controlKind string) *FormControlElement {
	if _, ok := formControlHandlers[controlKind]; !ok {
		log.Fatalf("Unknown form control kind: %s\n", controlKind)
		return nil
	}

	return &FormControlElement{
		props:       mttools.NewValues(),
		controlKind: controlKind,
	}
}

func (e *FormControlElement) GetTags() dhtml.TagList {
	var out dhtml.HtmlPiece

	if handler, ok := formControlHandlers[e.controlKind]; !ok {
		out.Append(handler.RenderF(e.props))
	} else {
		out.Append(dhtml.Dbg("Unknown form control kind: %s", e.controlKind))
	}

	return out.GetTags()
}
