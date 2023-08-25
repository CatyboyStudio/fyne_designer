package workspace

import (
	"noc"
	"time"
)

type WorkspaceProgress interface {
	ShowDelay() time.Duration
	Show()
	Close()
}

type WorkspaceExecutor func(w *Workspace) error

var Current *noc.NodeHost

func NewWorkspace() *noc.NodeHost {
	ws := newWorkspace()
	host := noc.NewHost(ws.Node())
	host.Run()
	Current = host
	return host
}

func ExecuteNodeTask(f noc.NodeExecutor, wp WorkspaceProgress) bool {
	this := Current
	if wp == nil {
		return this.Post(f)
	}
	ret := make(chan any)
	b := this.Post(func(w *noc.Node) error {
		defer func() {
			close(ret)
		}()
		return f(w)
	})
	if !b {
		return false
	}
	go func() {
		tm := time.NewTimer(wp.ShowDelay())
		defer tm.Stop()
		select {
		case <-ret:
			return
		case <-tm.C:
			wp.Show()
		}
		defer wp.Close()
		<-ret
	}()
	return b
}

func AddWorkspaceListener(lis func(WSEvent), cb func(int)) bool {
	return Current.Post(func(n *noc.Node) error {
		if n == nil {
			return nil
		}
		w := n.MainData.(*Workspace)
		w.lisid += 1
		w.listeners[w.lisid] = lis
		if cb != nil {
			cb(w.lisid)
		}
		return nil
	})
}

func RemoveWorkspaceListener(id int) {
	Current.Post(func(n *noc.Node) error {
		if n == nil {
			return nil
		}
		w := n.MainData.(*Workspace)
		delete(w.listeners, id)
		return nil
	})
}

func ExecWorkspaceTask(f WorkspaceExecutor, wp WorkspaceProgress) bool {
	return ExecuteNodeTask(func(n *noc.Node) error {
		var w *Workspace
		if n != nil {
			w = n.MainData.(*Workspace)
		}
		return f(w)
	}, wp)
}
