package designer_window

import "fyne_designer/workspace"

func (this *DesignerWindow) onWorkspaceEvent(ev workspace.WSEvent) {
	if ev.Event == workspace.EVENT_DOC_OPEN {
		this.toolp.docView.addDocument(ev.Data.(*workspace.Document))
	}
}
