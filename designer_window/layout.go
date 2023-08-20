package designer_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*messageBoxLayout)(nil)

type messageBoxLayout struct {
	anchor fyne.Position // x: -1/left 0/center 1/right, y: -1/bottom 0/middle 1/right
}

// NewminLayout creates a new minLayout instance
func NewMessageBoxLayout(anchor fyne.Position) fyne.Layout {
	return &messageBoxLayout{
		anchor: anchor,
	}
}

// Layout is called to pack all child objects into a specified size.
// For minLayout this sets all children to the full size passed.
func (m *messageBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	msz := m.MinSize(objects)
	nsz := msz.Min(size)
	topLeft := fyne.NewPos(0, 0)
	if m.anchor.X > 0 {
		topLeft.X = size.Width - nsz.Width
	} else if m.anchor.X == 0 {
		topLeft.X = (size.Width - nsz.Width) / 2
	}
	if m.anchor.Y < 0 {
		topLeft.Y = size.Height - nsz.Height
	} else if m.anchor.Y == 0 {
		topLeft.Y = (size.Height - nsz.Height) / 2
	}

	for _, child := range objects {
		child.Resize(size)
		child.Move(topLeft)
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For minLayout this is determined simply as the MinSize of the largest child.
func (m *messageBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		minSize = minSize.Max(child.MinSize())
	}

	return minSize
}

func NewMessageBox(anchor fyne.Position, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(NewMessageBoxLayout(anchor), objects...)
}
