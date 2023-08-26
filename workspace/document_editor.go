package workspace

import (
	"context"
	"fyne_widget/inspector"

	"fyne.io/fyne/v2/widget"
)

var _ (inspector.Editor) = (*Document)(nil)

func (this *Document) CreateInspectorGUI(ctx context.Context, form *widget.Form, label string) error {
	form.Append("Test", widget.NewEntry())
	return nil
}
