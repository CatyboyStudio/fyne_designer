package main

import (
	"fyne_designer/designer_window"
	"fyne_designer/workspace"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/CatyboyStudio/goapp_commons"
)

func main() {
	goapp_commons.Init("config.toml", "log.toml")

	fyne_widgets.InitFont()
	defer os.Unsetenv("FYNE_FONT")

	os.Setenv("FYNE_THEME", "dark")
	myApp := app.NewWithID("FyneDesigner.CatyboyStudio")

	goapp_commons.SupportLangs = append(goapp_commons.SupportLangs,
		// goapp_commons.NewLangInfo("zh-CHS"),
		goapp_commons.NewLangInfo("zh"),
	)
	goapp_commons.InitI18N("i18n")

	func() {
		w := workspace.NewWorkspace()
		defer w.Close()

		mainWindow := designer_window.NewDesignerWindow()
		mainWindow.Show()

		myApp.Run()
	}()

	// time.Sleep(time.Second)
}
