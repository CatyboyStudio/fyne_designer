package designer_window

import (
	"fmt"
	"fyne_designer/workspace"
)

func (this *DesignerWindow) onWorkspaceEvent(ev workspace.WSEvent) {
	switch ev.Event {
	case workspace.EVENT_DOC_OPEN:
		this.toolp.docView.addDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_CLOSE:
		this.toolp.docView.removeDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_ACTIVE:
		fmt.Println(ev.Event, ev.Data)
	}
}
