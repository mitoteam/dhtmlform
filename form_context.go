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
