package designer_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/goapp_commons"
)

type toolPanel struct {
	content fyne.CanvasObject
}

func newToolPanel() *toolPanel {
	return &toolPanel{
		//
	}
}

func (this *toolPanel) getContent() fyne.CanvasObject {
	if this.content == nil {
		this.content = this.build()
	}
	return this.content
}

func (this *toolPanel) build() fyne.CanvasObject {
	M := goapp_commons.GetMessage
	tabs := container.NewAppTabs(
		container.NewTabItem(M("ToolPanel.UITree.Title"), widget.NewLabel("World!")),
		container.NewTabItem(M("ToolPanel.Workspace.Title"), widget.NewLabel("Hello")),
	)
	tabs.SetTabLocation(container.TabLocationBottom)
	return tabs
}
