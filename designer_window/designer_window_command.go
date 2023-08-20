package designer_window

import (
	"fyne_designer/workspace"
)

func (this *DesignerWindow) syncToggleValue() {
	this.toggleLeftItem.Checked = this.toggleLeft
	this.toggleRightItem.Checked = this.toggleRight
	this.toggleItem.Checked = this.toggle
	this.toggleRefresh()
}

func (this *DesignerWindow) toggleValue(p *bool) {
	*p = !*p
	this.syncToggleValue()
	this.build_Designer()
}

func (this *DesignerWindow) commandToggleView() {
	this.toggleValue(&this.toggle)
}

func (this *DesignerWindow) commandToggleToolPanel() {
	this.toggleValue(&this.toggleLeft)
}

func (this *DesignerWindow) commandToggleInspectorPanel() {
	this.toggleValue(&this.toggleRight)
}

func (this *DesignerWindow) commandNewDocument() {
	ExecWorkspaceTask(func(w *workspace.Workspace) error {
		doc := workspace.NewDocument()
		return w.OpenDocument(doc)
	})
}
