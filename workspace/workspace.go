package workspace

import "noc"

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

type Workspace struct {
	noc.BaseNode

	documents      map[string]*Document
	activeDocument *Document

	lisid     int
	listeners map[int]func(WSEvent)
}

func newWorkspace() *Workspace {
	o := &Workspace{
		documents: make(map[string]*Document),
		listeners: make(map[int]func(WSEvent)),
	}
	o.CreateBaseNode(o)
	return o
}

func (this *Workspace) RaiseEvent(ev string, data any) {
	e := WSEvent{
		Event: ev,
		Data:  data,
	}
	for _, lis := range this.listeners {
		lis(e)
	}
}

func (this *Workspace) NextEvent(ev string, data any, next WorkspaceExecutor) {
	e := WSEvent{
		Event: ev,
		Data:  data,
		Next:  next,
	}
	for _, lis := range this.listeners {
		lis(e)
	}
}

func (this *Workspace) GetActiveDocument() *Document {
	return this.activeDocument
}
