package designer_window

import (
	"fyne_designer/workspace"
)

func (dw *DesignerWindow) onWorkspaceEvent(ev workspace.WSEvent) {
	switch ev.Event {
	case workspace.EVENT_DOC_OPEN:
		dw.toolp.docView.addDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_CLOSE:
		dw.toolp.docView.removeDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_UPDATE:
		dw.toolp.docView.updateDocument(ev.Data.(*workspace.Document))
	case workspace.EVENT_DOC_ACTIVE:
		if ev.Data == nil {
			dw.inspector.Unbind(dw.docbid)
		} else {
			doc := ev.Data.(*workspace.Document)
			bid, err := dw.inspector.Bind(doc, "")
			if err != nil {
				dw.log.Error("show Inspector fail",
					"docId", doc.GetId(),
					"docTitle", doc.GetTitle(),
					"error", err)
			} else {
				dw.docbid = bid
			}
		}
	}
}
