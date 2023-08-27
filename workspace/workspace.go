package workspace

import (
	"cbsutil/collections"
	"noc"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
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
	EVENT_DOC_UPDATE   = "doc_update"
	EVENT_DOC_SAVEFILE = "doc_savefile" // next
)

type HandlerForWorkspaceEvent func(WSEvent)

type Workspace struct {
	node *noc.Node

	dir string

	documents      map[string]*Document
	activeDocument *Document

	listeners collections.IdSlice[HandlerForWorkspaceEvent]
}

func newWorkspace() *Workspace {
	o := &Workspace{
		node:      noc.NewNode(),
		documents: make(map[string]*Document),
	}
	o.node.MainData = o
	wd, _ := os.Getwd()
	o.dir = fyne.CurrentApp().Preferences().StringWithFallback("Workspace_dir", wd)
	return o
}

func (th *Workspace) Node() *noc.Node {
	return th.node
}

func (th *Workspace) Dir() string {
	return th.dir
}

func (th *Workspace) SetDir(p string) {
	v, err := filepath.Abs(p)
	if err != nil {
		return
	}
	if v == th.dir {
		return
	}
	th.dir = v
	fyne.CurrentApp().Preferences().SetString("Workspace_dir", v)
}

func (th *Workspace) RaiseEvent(ev string, data any) {
	e := WSEvent{
		Event: ev,
		Data:  data,
	}
	for _, lis := range th.listeners.Data {
		lis(e)
	}
}

func (th *Workspace) NextEvent(ev string, data any, next WorkspaceExecutor) {
	e := WSEvent{
		Event: ev,
		Data:  data,
		Next:  next,
	}
	for _, lis := range th.listeners.Data {
		lis(e)
	}
}

func (th *Workspace) GetActiveDocument() *Document {
	return th.activeDocument
}
