package designer_window

import (
	"fyne_designer/widgets"
	"fyne_designer/workspace"
	"goapp_commons"
	"goapp_fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog"
)

var MainWindow *DesignerWindow

type DesignerWindow struct {
	window fyne.Window
	log    zerolog.Logger

	designerContainer *fyne.Container
	toolbar           *widget.Toolbar
	split3            *widgets.Split3

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
	loadErrId     int

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

func (this *DesignerWindow) shutdown() {
	MainWindow = nil
	this.window.SetOnClosed(nil)
	this.popupMessage.Close()
}

func (this *DesignerWindow) Show() {
	this.window.Resize(fyne.NewSize(740, 480))
	this.window.SetContent(this.build_Main())
	goapp_fyne.ShowMaximizeWindow(this.window)
	this.popupMessage.Start()
	workspace.AddWorkspaceListener(this.onWorkspaceEvent, func(id int) {
		this.lisid = id
	})
}

func (this *DesignerWindow) Close() {
	this.window.Close()
}

func ShowPopupError(err error) {
	goapp_commons.DefaultLogger().Warn().Err(err).Stack().Msg("ShowPopupError")
	if MainWindow != nil {
		MainWindow.popupMessage.AddErrorMessage(err.Error(), 10)
	}
}
