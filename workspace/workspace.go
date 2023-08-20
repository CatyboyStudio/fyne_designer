package workspace

type Workspace struct {
	r *WorkspaceHost

	lisid     int
	listeners map[int]func(any)
}

func newWorkspace(r *WorkspaceHost) *Workspace {
	return &Workspace{
		r:         r,
		listeners: make(map[int]func(any)),
	}
}

func (this *Workspace) RaiseEvent(ev any) {
	for _, lis := range this.listeners {
		lis(ev)
	}
}
