package dhtmlform

import (
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
	controlsData map[string]FormControlData

	rebuild     bool   // flag to rebuild form with same data again
	redirectUrl string // issue an redirect to this URL after processing
}

func NewFormData() *FormData {
	return &FormData{
		build_id:     "fd_" + mttools.RandomString(64),
		args:         mttools.NewValues(),
		params:       mttools.NewValues(),
		controlsData: make(map[string]FormControlData),
	}
}
