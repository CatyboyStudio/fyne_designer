package designer_window

import (
	"fyne_designer/widgets"
	"fyne_designer/workspace"
	"fyne_widget/inspector"
	"goapp_commons"
	"goapp_fyne"
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var MainWindow *DesignerWindow

type DesignerWindow struct {
	window fyne.Window
	log    *slog.Logger

	designerContainer *fyne.Container
	toolbar           *widget.Toolbar
	split3            *widgets.Split3

	inspector *inspector.Inspector
	docbid    int

	toggle          bool
	toggleLeft      bool
	toggleRight     bool
	toggleItem      *fyne.MenuItem
	toggleLeftItem  *fyne.MenuItem
	toggleRightItem *fyne.MenuItem
	toggleRefresh   func()

	toolp *toolPanel

	popupMessage  *goapp_fyne.PopupMessageManager
	messageBox    *fyne.Container
	messageShower fyne.CanvasObject
	// loadErrId     int

	lisid int
}

func NewDesignerWindow() *DesignerWindow {
	title := goapp_commons.GetMessage("MainWindow.Title")
	win := fyne.CurrentApp().NewWindow(title)
	o := &DesignerWindow{
		window:      win,
		toggleLeft:  true,
		toggleRight: true,
		log:         goapp_commons.NewLog("MainWindow"),
	}
	o.popupMessage = goapp_fyne.NewPopupMessageManager()
	o.window.SetOnClosed(o.shutdown)
	MainWindow = o
	win.SetMainMenu(o.build_MainMenu())
	win.SetMaster()
	return o
}

func (dw *DesignerWindow) shutdown() {
	MainWindow = nil
	dw.window.SetOnClosed(nil)
	dw.popupMessage.Close()
}

func (dw *DesignerWindow) Show() {
	dw.window.Resize(fyne.NewSize(740, 480))
	dw.window.SetContent(dw.build_Main())
	goapp_fyne.ShowMaximizeWindow(dw.window)
	dw.popupMessage.Start()
	workspace.AddWorkspaceListener(dw.onWorkspaceEvent, func(id int) {
		dw.lisid = id
	})
}

func (dw *DesignerWindow) Close() {
	dw.window.Close()
}

func ShowPopupError(err error) {
	slog.Warn("ShowPopupError", "error", err)
	if MainWindow != nil {
		MainWindow.popupMessage.AddErrorMessage(err.Error(), 10)
	}
}
