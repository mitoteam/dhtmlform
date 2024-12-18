package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

// Form control data to be stored in FormData between builds
type FormControlData struct {
	controlKind string
	value       any
}

// creates a copy of FormControlData and returns its pointer
func (fcd *FormControlData) getCopyPtr() *FormControlData {
	new_fcd := *fcd //simple value copy until we have primitives only in it

	return &new_fcd
}

type FormControlElement struct {
	id    string
	name  string
	label dhtml.HtmlPiece
	note  dhtml.HtmlPiece

	//additional properties
	props mttools.Values

	data FormControlData
}

var _ dhtml.ElementI = (*FormControlElement)(nil)

func NewFormControl(controlKind string, name string) *FormControlElement {
	if _, ok := GetFormControlHandler(controlKind); !ok {
		return nil
	}

	return &FormControlElement{
		name:  name,
		id:    dhtml.SafeId("id_" + controlKind + "_" + name),
		props: mttools.NewValues(),

		data: FormControlData{
			controlKind: controlKind,
		},
	}
}

func (e *FormControlElement) Default(v any) *FormControlElement {
	e.data.value = v
	return e
}

func (e *FormControlElement) Label(v any) *FormControlElement {
	e.label.Append(v)
	return e
}

func (e *FormControlElement) Note(v any) *FormControlElement {
	e.note.Append(v)
	return e
}

func (e *FormControlElement) GetId() string {
	return e.id
}

func (e *FormControlElement) GetTags() dhtml.TagList {
	var out dhtml.HtmlPiece

	if handler, ok := GetFormControlHandler(e.data.controlKind); ok {
		out.Append(handler.RenderF(e))
	} else {
		out.Append(dhtml.Dbg("Unknown form control kind: %s", e.data.controlKind))
	}

	return out.GetTags()
}

func (e *FormControlElement) renderLabel() *dhtml.LabelElement {
	return dhtml.NewLabel().For(e.GetId()).Styles("font-weight: bolder; vertical-align: top").
		Append(e.label)
}

func (e *FormControlElement) renderNote() *dhtml.Tag {
	return dhtml.Div().Styles("color: grey; font-size: 85%;").Append(e.note)
}
