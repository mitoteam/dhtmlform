package dhtmlform

import (
	"fmt"
	"strings"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

type FormControlHandler interface {
	RenderF(props mttools.Values) *dhtml.HtmlPiece
}

var formControlHandlers map[string]FormControlHandler

func init() {
	formControlHandlers = make(map[string]FormControlHandler)
}

func RegisterFormControlHandler(controlKind string, handler FormControlHandler) {
	controlKind = strings.TrimSpace(controlKind)

	if controlKind == "" {
		panic("controlKind should not be empty")
	}

	if handler == nil {
		panic("handler should not be nil")
	}

	if _, ok := formControlHandlers[controlKind]; ok {
		panic(fmt.Sprintf("handler for '%s' already registered", controlKind))
	}

	formControlHandlers[controlKind] = handler
}
