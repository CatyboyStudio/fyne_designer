package designer_window

import (
	"fyne.io/fyne/v2"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/CatyboyStudio/goapp_commons"
)

var MainWindow *DesignerWindow

type DesignerWindow struct {
	window fyne.Window

	popupMessage  *fyne_widgets.PopupMessageManager
	messageBox    *fyne.Container
	messageShower fyne.CanvasObject
	loadErrId     int
}

func NewDesignerWindow() *DesignerWindow {
	title := goapp_commons.GetMessage("MainWindow.Title")
	win := fyne.CurrentApp().NewWindow(title)
	win.SetMaster()
	o := &DesignerWindow{
		window: win,
	}
	o.popupMessage = fyne_widgets.NewPopupMessageManager()
	o.window.SetOnClosed(o.shutdown)
	MainWindow = o
	return o
}

func (this *DesignerWindow) shutdown() {
	MainWindow = nil
	this.window.SetOnClosed(nil)
	this.popupMessage.Close()
}

func (this *DesignerWindow) Show() {
	this.window.Resize(fyne.NewSize(740, 480))
	this.window.Show()
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
