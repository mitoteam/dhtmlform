package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

// Represents form state between builds from rendering through validation and submit.
// Constructed on first render from FormContext, stored on server side.
// Linked to browser representation by build_id.
type FormData struct {
	build_id string

	params mttools.Values // copied from FormContext each time form is rendered (even if it is being rebuild)
	args   mttools.Values // copied from FormContext on first build only and stored between builds

	//list of all form controls data: control name => Control element
	controlsData map[string]*FormControlData

	rebuild     bool   // flag to rebuild form with same data again
	redirectUrl string // issue an redirect to this URL after processing
}

func NewFormData() *FormData {
	return &FormData{
		build_id:     "fd_" + mttools.RandomString(64),
		args:         mttools.NewValues(),
		params:       mttools.NewValues(),
		controlsData: make(map[string]*FormControlData),
	}
}

func (fd *FormData) GetArg(name string) any {
	return fd.args.Get(name)
}

func (fd *FormData) GetParam(name string) any {
	return fd.params.Get(name)
}

func (fd *FormData) GetValue(name string) any {
	if controlDataPtr, ok := fd.controlsData[name]; ok {
		return controlDataPtr.value
	}

	return nil
}

func (fd *FormData) SetValue(name string, v any) *FormData {
	if controlDataPtr, ok := fd.controlsData[name]; ok {
		controlDataPtr.value = v
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

// Walker function to set control values from FormData if form being rebuild after post
func (fd *FormData) processControlDataWalkerF(e dhtml.ElementI, args ...any) {
	if control, ok := e.(*FormControlElement); ok {
		if storedControlDataPtr, ok := fd.controlsData[control.name]; ok {
			if storedControlDataPtr.controlKind == control.data.controlKind {
				control.data = *storedControlDataPtr
			} else {
				// kind does not match, no need to store it at all
				// https://stackoverflow.com/questions/23229975/is-it-safe-to-remove-selected-keys-from-map-within-a-range-loop
				delete(fd.controlsData, control.name)
			}
		} else {
			fd.controlsData[control.name] = control.data.getCopyPtr()
		}

		//control.Note(fmt.Sprintf("DBG: processControlDataWalkerF Walked, rebuild: %t", fd.rebuild))
	}
}
