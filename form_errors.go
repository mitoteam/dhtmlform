package dhtmlform

import (
	"github.com/mitoteam/dhtml"
)

type FormControlErrorT struct {
	label  dhtml.HtmlPiece   //form control label to render
	errors []dhtml.HtmlPiece // errors list
}

// List of form control errors (name => errors list, empty string name for common errors)
type FormErrorsT map[string]FormControlErrorT
