package designer_window

import (
	"cbsutil/logi"
	"fyne_designer/workspace"
)

func (this *DesignerWindow) onWorkspaceEvent(ev workspace.WSEvent) {
	switch ev.Event {
	case workspace.EVENT_DOC_OPEN:
		this.toolp.docView.addDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_CLOSE:
		this.toolp.docView.removeDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_ACTIVE:
		logi.Debug.Printf("document active: %v", ev.Data)
	}
}
