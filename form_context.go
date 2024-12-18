package dhtmlform

import (
	"net/http"

	"github.com/mitoteam/mttools"
)

// Initial data to build a form
type FormContext struct {
	params      mttools.Values // copied to form data each time form is rendered (even if it is being rebuild)
	args        mttools.Values // copied to form data on first build only and stored between builds
	redirectUrl string         // issue an redirect to this URL (FormData's redirectUrl has priority)

	w http.ResponseWriter
	r *http.Request
}

func NewFormContext(w http.ResponseWriter, r *http.Request) *FormContext {
	fc := &FormContext{
		w: w, r: r,
		params: mttools.NewValues(),
		args:   mttools.NewValues(),
	}

	return fc
}

func (fc *FormContext) SetParam(key string, v any) *FormContext {
	fc.params.Set(key, v)
	return fc
}

func (fc *FormContext) GetParam(key string) any {
	return fc.params.Get(key)
}

func (fc *FormContext) SetArg(key string, v any) *FormContext {
	fc.args.Set(key, v)
	return fc
}

func (fc *FormContext) GetArg(key string) any {
	return fc.args.Get(key)
}

func (fc *FormContext) SetRedirect(url string) *FormContext {
	fc.redirectUrl = url
	return fc
}

// Check build_id in POST values, try to find form data in data store, re-hydrate it with submitted values and return.
// Returns nil if there is no FormData to rebuild.
func (fc *FormContext) formDataFromPOST() (fd *FormData) {
	build_id := fc.r.PostFormValue(hiddenBuildIdFieldName)

	// check if it is being re-build (from POST request)
	if build_id == "" {
		return nil
	}

	// check if build_id is in store
	fd, ok := formDataStore[build_id]
	if !ok {
		return nil
	}

	//refresh params from context
	fd.params.CopyFrom(fc.params)

	//re-hydrate form_data.values from POST data
	for name, controlDataPtr := range fd.controlsData {
		controlDataPtr.value = nil // empty by default

		if rawPostValue, ok := fc.r.PostForm[name]; ok {
			if len(rawPostValue) == 1 { //array of single element, just take first one
				controlDataPtr.value = rawPostValue[0]
			} else {
				controlDataPtr.value = rawPostValue
			}
		}

		if handler, ok := GetFormControlHandler(controlDataPtr.controlKind); ok {
			if handler.ProcessPostValueF != nil {
				controlDataPtr.value = handler.ProcessPostValueF(controlDataPtr.value)
			}
		}
	}

	return fd
}
