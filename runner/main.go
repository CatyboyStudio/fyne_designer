package main

import (
	"fyne_designer/designer_window"
	"fyne_designer/workspace"
	"os"

	"goapp_commons"
	"goapp_fyne"

	"fyne.io/fyne/v2/app"
)

func main() {
	goapp_commons.Init("config.toml", "log.toml")

	goapp_fyne.InitFont()
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
		// testui()

		myApp.Run()
	}()
}
