package designer_window

import (
	"fyne_designer/workspace"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/goapp_commons"
)

type wsProgress struct {
	count   int32
	d       dialog.Dialog
	content *widget.ProgressBarInfinite
}

// Close implements workspace.WorkspaceProgress.
func (this *wsProgress) Close() {
	v := atomic.LoadInt32(&this.count)
	if v > 0 {
		if atomic.CompareAndSwapInt32(&this.count, v, v-1) {
			if v == 1 {
				this.doClose()
			}
		}
	}
}

// Show implements workspace.WorkspaceProgress.
func (this *wsProgress) Show() {
	v := atomic.LoadInt32(&this.count)
	atomic.CompareAndSwapInt32(&this.count, v, v+1)
	if v == 0 {
		this.doShow()
	}
}

// ShowDelay implements workspace.WorkspaceProgress.
func (*wsProgress) ShowDelay() time.Duration {
	return time.Second * 1
}

func (this *wsProgress) doShow() {
	if this.d == nil {
		w := MainWindow.window
		M := goapp_commons.GetMessage

		if this.content == nil {
			this.content = widget.NewProgressBarInfinite()
		}
		d := dialog.NewCustom(M("MainWindow.Progress.Title"), M("MainWindow.Progress.Close"), this.content, w)
		d.Resize(fyne.NewSize(200, 60))
		d.SetOnClosed(func() {
			this.content.Stop()
			this.d = nil
		})
		d.Show()
		this.d = d
	}
	this.content.Start()
}

func (this *wsProgress) doClose() {
	if this.d != nil {
		this.d.Hide()
		this.content.Stop()
	}
}

var inswsProgress *wsProgress = &wsProgress{}

func ExecWorkspaceTask(f workspace.WorkspaceExecutor) bool {
	return workspace.ExecWorkspaceTask(f, inswsProgress)
}
