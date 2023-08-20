package workspace

import (
	"fyne.io/fyne/v2"
	"github.com/dave/jennifer/jen"
)

type ElementPrototype interface {
	TypeName() string

	Kind() ElementKind

	ToJson(e *Element) map[string]any

	FromJson(e *Element, data map[string]any)

	BuildCanvasObject(e *Element) fyne.CanvasObject

	BuildCode(e *Element) jen.Code
}
