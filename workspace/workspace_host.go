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

type WorkspaceExecutor func(w *Workspace) (any, error)

var Current *noc.NodeHost

func NewWorkspace() *noc.NodeHost {
	ws := newWorkspace()
	host := noc.NewHost(ws.Node())
	host.Run()
	Current = host
	return host
}

func PostNodeTask(f noc.NodeExecutor, wp WorkspaceProgress) bool {
	this := Current
	if wp == nil {
		return this.Post(f)
	}
	ret := make(chan int)
	b := this.Post(func(w *noc.Node) (any, error) {
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

func AddWorkspaceListener(lis HandlerForWorkspaceEvent, cb func(int)) bool {
	return Current.Post(func(n *noc.Node) (any, error) {
		if n == nil {
			return nil, nil
		}
		w := n.MainData.(*Workspace)
		id := w.listeners.Add(lis)
		if cb != nil {
			cb(id)
		}
		return nil, nil
	})
}

func RemoveWorkspaceListener(id int) {
	Current.Post(func(n *noc.Node) (any, error) {
		if n == nil {
			return nil, nil
		}
		w := n.MainData.(*Workspace)
		w.listeners.Remove(id)
		return nil, nil
	})
}

func PostWorkspaceTask(f WorkspaceExecutor, wp WorkspaceProgress) bool {
	return PostNodeTask(func(n *noc.Node) (any, error) {
		var w *Workspace
		if n != nil {
			w = n.MainData.(*Workspace)
		}
		return f(w)
	}, wp)
}

func ExecuteWorkspaceTask(f WorkspaceExecutor) (any, error) {
	return Current.Execute(func(n *noc.Node) (any, error) {
		var w *Workspace
		if n != nil {
			w = n.MainData.(*Workspace)
		}
		return f(w)
	})
}
