package dhtmlform

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitoteam/dhtml"
)

type FormControlHandler struct {
	// [required] renders control
	RenderF func(control *FormControlElement) dhtml.HtmlPiece

	// [optional] preprocesses value from POST values
	ProcessPostValueF func(rawValue any) any
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

	if handler.RenderF == nil {
		panic("handler.RenderF is not set")
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
