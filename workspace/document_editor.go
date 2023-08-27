package workspace

import (
	"fyne_widget/inspector"

	"fyne.io/fyne/v2/widget"
)

var _ (inspector.EditorBuilder) = (*Document)(nil)

// BuildEditor implements inspector.EditorBuilder.
func (th *Document) BuildEditor(ed *inspector.Editor) error {
	w := widget.NewEntry()
	w.Text = ""
	ed.Form.Append("Test", w)
	return nil
}
