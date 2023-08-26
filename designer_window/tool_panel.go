package designer_window

import (
	"goapp_commons"
	"goapp_fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

func (tp *toolPanel) build() fyne.CanvasObject {
	if tp.content == nil {
		M := goapp_commons.GetMessage
		tabs := container.NewAppTabs(
			container.NewTabItem(M("ToolPanel.Workspace.Title"), tp.build_Workspae()),
			container.NewTabItem(M("ToolPanel.UITree.Title"), widget.NewLabel("World!")),
		)
		tabs.SetTabLocation(container.TabLocationBottom)
		tp.content = tabs
	}
	return tp.content
}

func (tp *toolPanel) build_Workspae() fyne.CanvasObject {
	M := goapp_commons.GetMessage
	ac := goapp_fyne.NewAccordionBox(
		widget.NewAccordionItem(M("ToolPanel.Workspace.Document.Title"), tp.build_Workspace_DocumentView()),
		widget.NewAccordionItem("B", widget.NewLabel("Two")),
	)
	ac.Items[0].Open = true
	return ac
}

func (tp *toolPanel) build_Workspace_DocumentView() fyne.CanvasObject {
	view := NewDocumentView()
	tp.docView = view
	list := view.Build()
	return container.NewMax(list)
}
