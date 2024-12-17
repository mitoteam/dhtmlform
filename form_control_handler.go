package dhtmlform

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
)

type FormControlHandler struct {
	RenderF func(props mttools.Values) dhtml.HtmlPiece
}

var formControlHandlers map[string]*FormControlHandler

func init() {
	formControlHandlers = make(map[string]*FormControlHandler)
}

func RegisterFormControlHandler(controlKind string, handler *FormControlHandler) {
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

func GetFormControlHandler(controlKind string) (*FormControlHandler, bool) {
	if handler, ok := formControlHandlers[controlKind]; ok {
		return handler, true
	}

	log.Fatalf("Unknown form control kind: %s\n", controlKind)
	return nil, false
}
