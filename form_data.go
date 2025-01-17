package dhtmlform

import (
	"time"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

// Represents form state between builds from rendering through validation and submit.
// Constructed on first render from FormContext, stored on server side.
// Linked to browser representation by build_id.
type FormData struct {
	build_id string    // unique string to identify form build
	created  time.Time // time of form data creation (to be able to expire unused ones periodically)

	params mttools.Values // copied from FormContext each time form is rendered (even if it is being rebuild)
	args   mttools.Values // copied from FormContext on first build only and stored between builds

	//list of all form controls data: control name => Control element
	controlsData formControlsDataT

	// form controls errors (if any werer set during validation)
	errors FormErrors

	rebuild     bool   // flag to rebuild form with same data again
	redirectUrl string // issue an redirect to this URL after processing
}

type formControlsDataT map[string]*FormControlData

func NewFormData() *FormData {
	return &FormData{
		build_id:     "fd_" + mttools.RandomString(64),
		created:      time.Now(),
		args:         mttools.NewValues(),
		params:       mttools.NewValues(),
		controlsData: make(formControlsDataT),
		errors:       make(FormErrors),
	}
}

func (fd *FormData) GetArg(name string) any {
	return fd.args.Get(name)
}

func (fd *FormData) GetParam(name string) any {
	return fd.params.Get(name)
}

func (fd *FormData) SetParam(name string, value any) {
	fd.params.Set(name, value)
}

// Shorthand for GetControlValue(name)
func (fd *FormData) GetValue(name string) any {
	return fd.GetControlValue(name)
}

func (fd *FormData) GetControlValue(name string) any {
	if controlDataPtr, ok := fd.controlsData[name]; ok {
		return controlDataPtr.Value
	}

	return nil
}

func (fd *FormData) SetControlValue(name string, v any) *FormData {
	if controlDataPtr, ok := fd.controlsData[name]; ok {
		controlDataPtr.Value = v
	}

	return fd
}

func (fd *FormData) IsRebuild() bool {
	return fd.rebuild
}

func (fd *FormData) SetRebuild(v bool) {
	fd.rebuild = v
}

func (fd *FormData) GetRedirect() string {
	return fd.redirectUrl
}

func (fd *FormData) SetRedirect(url string) {
	fd.redirectUrl = url
}

func (fd *FormData) SetError(controlName string, v any) {
	var controlErrors *FormControlErrors
	var ok bool

	if controlErrors, ok = fd.errors[controlName]; !ok {
		//no errors for this form control added yet
		controlErrors = &FormControlErrors{}

		if controlDataPtr, ok := fd.controlsData[controlName]; ok {
			controlErrors.Label = controlDataPtr.label
			controlDataPtr.isError = true
		}

		fd.errors[controlName] = controlErrors
	}

	controlErrors.Errors = append(controlErrors.Errors, *dhtml.Piece(v))
}

func (fd *FormData) HasError() bool {
	return len(fd.errors) > 0
}

func (fd *FormData) ClearErrors() {
	fd.errors = make(FormErrors) //new empty map

	//clear controls error flags
	for _, controlDataPtr := range fd.controlsData {
		controlDataPtr.isError = false
	}
}

// some common and very basic validations
func (fd *FormData) validateFormControls() {
	for controlName, controlDataPtr := range fd.controlsData {
		if controlDataPtr.isRequired {
			if mttools.IsEmpty(controlDataPtr.Value) {
				fd.SetError(controlName, "value is required")
			}
		}
	}
}

// Walker function to set control values from FormData if form being rebuild after post
func (fd *FormData) processControlDataWalkerF(e dhtml.ElementI, args ...any) {
	control, ok := e.(FormControlElementI)

	if !ok {
		return
	}

	controlData := control.GetControlData()

	if storedControlDataPtr, ok := fd.controlsData[control.GetName()]; ok {
		//check control kind
		if storedControlDataPtr.controlKind == controlData.controlKind {
			storedControlDataPtr.label = controlData.label //update label if changed since last build
			control.SetControlData(storedControlDataPtr)
		} else {
			// kind does not match, no need to store it at all
			delete(fd.controlsData, control.GetName())
		}
	} else {
		//new control, add new control data to store for it
		fd.controlsData[control.GetName()] = controlData.getCopyPtr()
	}

	//control.Note(fmt.Sprintf("DBG: processControlDataWalkerF Walked, rebuild: %t", fd.rebuild))
}
