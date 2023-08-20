package workspace

import (
	"fyne_designer/widgets"
)

type ElementKind int

const (
	ElementKindValue = iota
	ElementKindFactory
	ElementKindObject
	ElementKindProperty
)

type Element struct {
	Id        int
	Kind      ElementKind
	Value     string
	Prototype ElementPrototype

	Cell *widgets.DesignCellWidget
}

func NewElementValue(id int, val string, pro ElementPrototype) *Element {
	return &Element{
		Id:        id,
		Value:     val,
		Prototype: pro,
	}
}
