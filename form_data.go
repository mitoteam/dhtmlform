package dhtmlform

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

type FormErrorsT map[string][]dhtml.HtmlPiece

type FormData struct {
	build_id string

	args   mttools.Values
	params mttools.Values
	values mttools.Values

	labels dhtml.NamedHtmlPieces

	errorList   FormErrorsT //map of error lists by form item name
	rebuild     bool        // rebuild form with same data again
	redirectUrl string      // issue an redirect to this URL
}
