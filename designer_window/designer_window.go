package designer_window

import (
	"fyne_designer/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/CatyboyStudio/goapp_commons"
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

	popupMessage  *fyne_widgets.PopupMessageManager
	messageBox    *fyne.Container
	messageShower fyne.CanvasObject
	loadErrId     int
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
	o.popupMessage = fyne_widgets.NewPopupMessageManager()
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
	fyne_widgets.ShowMaximizeWindow(this.window)
	this.popupMessage.Start()
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
