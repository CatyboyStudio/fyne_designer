package designer_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/CatyboyStudio/goapp_commons"
)

type toolPanel struct {
	content fyne.CanvasObject

	docView *DocumentView
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
		container.NewTabItem(M("ToolPanel.Workspace.Title"), this.build_Workspae()),
		container.NewTabItem(M("ToolPanel.UITree.Title"), widget.NewLabel("World!")),
	)
	tabs.SetTabLocation(container.TabLocationBottom)
	return tabs
}

func (this *toolPanel) build_Workspae() fyne.CanvasObject {
	M := goapp_commons.GetMessage
	ac := fyne_widgets.NewAccordionBox(
		widget.NewAccordionItem(M("ToolPanel.Workspace.Document.Title"), this.build_Workspace_DocumentView()),
		widget.NewAccordionItem("B", widget.NewLabel("Two")),
	)
	ac.Items[0].Open = true
	return ac
}

func (this *toolPanel) build_Workspace_DocumentView() fyne.CanvasObject {
	view := NewDocumentView()
	this.docView = view
	list := view.Build()
	return container.NewMax(list)
}
