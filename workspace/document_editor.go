package workspace

import "fyne.io/fyne/v2/widget"

func (this *Document) CreateInspectorGUI(form *widget.Form, label string) error {
	form.Append("Test", widget.NewEntry())
	return nil
}
