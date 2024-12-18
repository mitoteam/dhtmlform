package dhtmlform

import "github.com/mitoteam/dhtml"

const (
	errorBlockStyle = "border: 2px solid red; background-color: rgb(255 0 0 / 0.2);"
)

// package settings to be overridden
type settingsType struct {
	// function to render errors block if there are any after form validation
	FormErrorsRenderF func(fe *FormErrors) dhtml.HtmlPiece
}

var settings *settingsType

func Settings() *settingsType {
	return settings
}

// default implementations
func init() {
	settings = &settingsType{
		FormErrorsRenderF: func(fe *FormErrors) (out dhtml.HtmlPiece) {
			container := dhtml.Div().Styles(errorBlockStyle).Styles("margin-bottom: 10px; padding: 10px;")

			for _, controlErrors := range *fe {

				for _, controlError := range controlErrors.Errors {
					errorOut := dhtml.Div()

					if !controlErrors.Label.IsEmpty() {
						errorOut.
							Append(dhtml.Span().Styles("font-weight: bolder; color: red;").Append(controlErrors.Label)).
							Text(":")
					}

					errorOut.Append(controlError)

					container.Append(errorOut)
				}
			}

			out.Append(container)
			return out
		},
	}
}
