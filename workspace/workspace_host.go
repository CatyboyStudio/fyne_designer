package workspace

import (
	"goapp_commons"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
)

type WorkspaceProgress interface {
	ShowDelay() time.Duration
	Show()
	Close()
}

type WorkspaceExecutor func(w *Workspace) error

var Current *WorkspaceHost

func NewWorkspace() *WorkspaceHost {
	w := &WorkspaceHost{
		log:    goapp_commons.NewLog("Workspace"),
		closeC: make(chan any, 0),
		execs:  make(chan WorkspaceExecutor, 16),
	}
	w.w = newWorkspace(w)
	go w.run()
	Current = w
	return w
}

type WorkspaceHost struct {
	log zerolog.Logger
	w   *Workspace

	status int32
	closeC chan any
	execs  chan WorkspaceExecutor
}

func (this *WorkspaceHost) exec(f WorkspaceExecutor) bool {
	select {
	case <-this.closeC:
		return false
	case this.execs <- f:
		return true
	}
}

func (this *WorkspaceHost) doInit() error {
	return nil
}

func (this *WorkspaceHost) run() {
	if !atomic.CompareAndSwapInt32(&this.status, 0, 1) {
		return
	}
	go func() {
		askClose := false
		defer func() {
			this.log.Debug().Msg("shutdown")
			this.shutdown()
		}()
		this.log.Debug().Msg("run")

		err := this.doInit()
		if err != nil {
			this.log.Error().Err(err).Stack().Msg("init fail")
			this.Close()
		}

		for {
			var f WorkspaceExecutor
			if askClose {
				select {
				case f = <-this.execs:
					break
				default:
					return
				}
			} else {
				select {
				case <-this.closeC:
					this.log.Debug().Msg("ask close")
					askClose = true
				case f = <-this.execs:
					break
				}
			}
			if f != nil {
				this.log.Debug().Msg("exec")
				this._exec(f)
			}
		}
	}()
}

func (this *WorkspaceHost) _exec(f WorkspaceExecutor) {
	defer func() {
		erro := recover()
		if erro != nil {
			s := string(debug.Stack())
			if err, ok := erro.(error); ok {
				this.log.Error().Err(err).Msgf("exec error: %s", s)
			} else {
				this.log.Error().Msgf("exec error: %v, %s", erro, s)
			}
		}
	}()
	err := f(this.w)
	if err != nil {
		this.log.Error().Err(err).Stack().Msg("exec fail")
	}
}

func (this *WorkspaceHost) shutdown() {
	atomic.StoreInt32(&this.status, -2)
}

func (this *WorkspaceHost) Exec(f WorkspaceExecutor, wp WorkspaceProgress) bool {
	if wp == nil {
		return this.Post(f)
	}
	ret := make(chan any)
	b := this.exec(func(w *Workspace) error {
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

func (this *WorkspaceHost) Post(f WorkspaceExecutor) bool {
	return this.exec(f)
}

func (this *WorkspaceHost) Close() {
	if atomic.CompareAndSwapInt32(&this.status, 1, -1) {
		close(this.closeC)
	}
}

func (this *WorkspaceHost) AddListener(lis func(WSEvent), cb func(int)) bool {
	return this.Post(func(w *Workspace) error {
		w.lisid += 1
		w.listeners[w.lisid] = lis
		if cb != nil {
			cb(w.lisid)
		}
		return nil
	})
}

func (this *WorkspaceHost) RemoveListener(id int) {
	this.Post(func(w *Workspace) error {
		delete(w.listeners, id)
		return nil
	})
}

func ExecWorkspaceTask(f WorkspaceExecutor, wp WorkspaceProgress) bool {
	c := Current
	if c == nil {
		return false
	}
	return c.Exec(func(w *Workspace) error {
		return f(w)
	}, wp)
}
