package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

// Form control data to be stored in FormData between builds
type FormControlData struct {
	controlKind string
	label       dhtml.HtmlPiece //how to refer eto this control when saying about it to user (in error message for example)
	isRequired  bool            // value should be set

	isError bool //flag indicating control has some errors from validation (to be used in renderer)

	props mttools.Values //custom properties
	Value any
}

// creates a copy of FormControlData and returns its pointer
func (fcd *FormControlData) getCopyPtr() *FormControlData {
	new_fcd := *fcd

	new_fcd.props.CopyFrom(&fcd.props)

	return &new_fcd
}

func (fcd *FormControlData) GetProp(key string) any {
	return fcd.props.Get(key)
}

func (fcd *FormControlData) HasProp(key string) bool {
	_, ok := fcd.props.GetOk(key)
	return ok
}

func (fcd *FormControlData) SetProp(key string, value any) {
	fcd.props.Set(key, value)
}

type FormControlElementI interface {
	dhtml.ElementI

	GetName() string
	GetControlData() *FormControlData
	SetControlData(*FormControlData)
}

// default implementation of FormControlElementI
type FormControlElement struct {
	id          string
	name        string
	placeholder string
	note        dhtml.HtmlPiece
	props       mttools.Values //custom properties

	data FormControlData
}

// force interface implementation
var _ FormControlElementI = (*FormControlElement)(nil)

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
			props:       mttools.NewValues(),
		},
	}
}

func (e *FormControlElement) GetName() string {
	return e.name
}

func (e *FormControlElement) GetControlData() *FormControlData {
	return &e.data
}

func (e *FormControlElement) SetControlData(data *FormControlData) {
	e.data = *data
}

// Sets default control value. Used only for initial form build. Replaced by POST value when re-builds.
func (e *FormControlElement) Default(v any) *FormControlElement {
	e.data.Value = v
	return e
}

func (e *FormControlElement) GetValue() any {
	return e.data.Value
}

func (e *FormControlElement) Label(v any) *FormControlElement {
	e.data.label.Append(v)
	return e
}

func (e *FormControlElement) GetLabel() *dhtml.HtmlPiece {
	return &e.data.label
}

func (e *FormControlElement) IsError() bool {
	return e.data.isError
}

func (e *FormControlElement) SetRequired(b bool) *FormControlElement {
	e.data.isRequired = b
	return e
}

// Shorthand for SetRequired(true)
func (e *FormControlElement) Require() *FormControlElement {
	e.data.isRequired = true
	return e
}

func (e *FormControlElement) IsRequired() bool {
	return e.data.isRequired
}

func (e *FormControlElement) Note(v any) *FormControlElement {
	e.note.Append(v)
	return e
}

func (e *FormControlElement) GetNote() *dhtml.HtmlPiece {
	return &e.note
}

func (e *FormControlElement) Placeholder(s string) *FormControlElement {
	e.placeholder = s
	return e
}

func (e *FormControlElement) GetPlaceholder() string {
	return e.placeholder
}

func (e *FormControlElement) GetId() string {
	return e.id
}

func (e *FormControlElement) GetProp(key string) any {
	return e.props.Get(key)
}

func (e *FormControlElement) SetProp(key string, value any) *FormControlElement {
	e.props.Set(key, value)
	return e
}

func (e *FormControlElement) HasProp(key string) bool {
	return e.props.Has(key)
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
	labelElement := dhtml.NewLabel().For(e.GetId()).Styles("font-weight: bolder; vertical-align: top;").
		Append(e.GetLabel())

	if e.IsRequired() {
		labelElement.Append(dhtml.Span().Styles("color: red;").Text("*"))
	}

	return labelElement
}

func (e *FormControlElement) renderNote() *dhtml.Tag {
	return dhtml.Div().Styles("color: grey; font-size: 85%;").Append(e.note)
}
