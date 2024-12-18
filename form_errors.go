package dhtmlform

import (
	"github.com/mitoteam/dhtml"
)

type FormControlErrors struct {
	Label  dhtml.HtmlPiece   //form control label to render
	Errors []dhtml.HtmlPiece // errors list
}

// List of form control errors (name => errors list, empty string name for common errors)
type FormErrors map[string]*FormControlErrors
