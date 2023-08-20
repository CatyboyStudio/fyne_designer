package workspace

type WSEvent struct {
	Event string
	Data  any
}

const (
	EVENT_DOC_OPEN  = "doc_open"
	EVENT_DOC_CLOSE = "doc_close"
)

type Workspace struct {
	r *WorkspaceHost

	documents map[string]*Document

	lisid     int
	listeners map[int]func(WSEvent)
}

func newWorkspace(r *WorkspaceHost) *Workspace {
	return &Workspace{
		r:         r,
		documents: make(map[string]*Document),
		listeners: make(map[int]func(WSEvent)),
	}
}

func (this *Workspace) RaiseEvent(ev string, data any) {
	e := WSEvent{ev, data}
	for _, lis := range this.listeners {
		lis(e)
	}
}
