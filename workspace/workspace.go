package workspace

type Workspace struct {
	r *WorkspaceHost
}

func newWorkspace(r *WorkspaceHost) *Workspace {
	return &Workspace{
		r: r,
	}
}
