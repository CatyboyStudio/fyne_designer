package designer_window

import (
	"fyne_designer/workspace"
	"goapp_commons"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type wsProgress struct {
	count   int32
	d       dialog.Dialog
	content *widget.ProgressBarInfinite
}

// Close implements workspace.WorkspaceProgress.
func (wsp *wsProgress) Close() {
	v := atomic.LoadInt32(&wsp.count)
	if v > 0 {
		if atomic.CompareAndSwapInt32(&wsp.count, v, v-1) {
			if v == 1 {
				wsp.doClose()
			}
		}
	}
}

// Show implements workspace.WorkspaceProgress.
func (wsp *wsProgress) Show() {
	v := atomic.LoadInt32(&wsp.count)
	atomic.CompareAndSwapInt32(&wsp.count, v, v+1)
	if v == 0 {
		wsp.doShow()
	}
}

// ShowDelay implements workspace.WorkspaceProgress.
func (*wsProgress) ShowDelay() time.Duration {
	return time.Second * 1
}

func (wsp *wsProgress) doShow() {
	if wsp.d == nil {
		w := MainWindow.window
		M := goapp_commons.GetMessage

		if wsp.content == nil {
			wsp.content = widget.NewProgressBarInfinite()
		}
		d := dialog.NewCustom(M("MainWindow.Progress.Title"), M("MainWindow.Progress.Close"), wsp.content, w)
		d.Resize(fyne.NewSize(200, 60))
		d.SetOnClosed(func() {
			wsp.content.Stop()
			wsp.d = nil
		})
		d.Show()
		wsp.d = d
	}
	wsp.content.Start()
}

func (wsp *wsProgress) doClose() {
	if wsp.d != nil {
		wsp.d.Hide()
		wsp.content.Stop()
	}
}

var inswsProgress *wsProgress = &wsProgress{}

func InvokeWorkspaceTask(f workspace.WorkspaceExecutor) bool {
	return workspace.PostWorkspaceTask(f, inswsProgress)
}
