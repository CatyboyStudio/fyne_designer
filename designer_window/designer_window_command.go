package designer_window

import (
	"fyne_designer/workspace"
)

func (dw *DesignerWindow) syncToggleValue() {
	dw.toggleLeftItem.Checked = dw.toggleLeft
	dw.toggleRightItem.Checked = dw.toggleRight
	dw.toggleItem.Checked = dw.toggle
	dw.toggleRefresh()
}

func (dw *DesignerWindow) toggleValue(p *bool) {
	*p = !*p
	dw.syncToggleValue()
	dw.build_Designer()
}

func (dw *DesignerWindow) commandToggleView() {
	dw.toggleValue(&dw.toggle)
}

func (dw *DesignerWindow) commandToggleToolPanel() {
	dw.toggleValue(&dw.toggleLeft)
}

func (dw *DesignerWindow) commandToggleInspectorPanel() {
	dw.toggleValue(&dw.toggleRight)
}

func (dw *DesignerWindow) commandNewDocument() {
	ExecWorkspaceTask(func(w *workspace.Workspace) error {
		obj := w.Node().NewObject()
		com, err := obj.AddComponent(workspace.DOC_COMTYPE)
		if err != nil {
			w.Node().DeleteObject(obj)
			return err
		}
		return w.OpenDocument(com.(*workspace.Document))
	})
}
