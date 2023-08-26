package workspace

import (
	"cbsutil/collections"
	"noc"
)

type WSEvent struct {
	Event string
	Data  any
	Next  any
}

const (
	EVENT_DOC_OPEN     = "doc_open"
	EVENT_DOC_CLOSE    = "doc_close"
	EVENT_DOC_ACTIVE   = "doc_active"
	EVENT_DOC_SAVEFILE = "doc_savefile" // next
)

type HandlerForWorkspaceEvent func(WSEvent)

type Workspace struct {
	node *noc.Node

	documents      map[string]*Document
	activeDocument *Document

	listeners collections.IdSlice
}

func newWorkspace() *Workspace {
	o := &Workspace{
		node:      noc.NewNode(),
		documents: make(map[string]*Document),
	}
	o.node.MainData = o
	return o
}

func (this *Workspace) Node() *noc.Node {
	return this.node
}

func (this *Workspace) RaiseEvent(ev string, data any) {
	e := WSEvent{
		Event: ev,
		Data:  data,
	}
	for _, lis := range this.listeners.Data {
		lis.(HandlerForWorkspaceEvent)(e)
	}
}

func (this *Workspace) NextEvent(ev string, data any, next WorkspaceExecutor) {
	e := WSEvent{
		Event: ev,
		Data:  data,
		Next:  next,
	}
	for _, lis := range this.listeners.Data {
		lis.(HandlerForWorkspaceEvent)(e)
	}
}

func (this *Workspace) GetActiveDocument() *Document {
	return this.activeDocument
}
