package designer_window

import (
	"goapp_commons"
	"goapp_fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

const DESIGNER_SITE = "https://catyboy.itch.io/"
const DESIGNER_GITHUB = "https://github.com/CatyboyStudio/FyneDesigner"

func (dw *DesignerWindow) build_MainMenu() *fyne.MainMenu {
	a := fyne.CurrentApp()
	w := dw.window
	M := goapp_commons.GetMessage

	changeDirItem := fyne.NewMenuItem(M("MainMenu.Workspace.ChangeDir"), func() {
		dw.commandChangeDir()
	})

	quitItem := fyne.NewMenuItem(M("MainMenu.Workspace.Quit"), func() {
		a.Quit()
	})
	quitItem.IsQuit = true

	wsp := fyne.NewMenu(M("MainMenu.Workspace.Title"),
		changeDirItem,
		fyne.NewMenuItemSeparator(),
		quitItem,
	)

	newItem := fyne.NewMenuItem(M("MainMenu.File.NewDocument"), func() {
		dw.commandNewDocument()
	})

	// openSettings := func() {
	// 	w := a.NewWindow("Fyne Settings")
	// 	w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
	// 	w.Resize(fyne.NewSize(480, 480))
	// 	w.Show()
	// }
	// settingsItem := fyne.NewMenuItem("Settings", openSettings)
	// settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	// settingsItem.Shortcut = settingsShortcut
	// w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
	// 	openSettings()
	// })

	file := fyne.NewMenu(M("MainMenu.File.Title"),
		newItem,
	)

	performToggle := dw.commandToggleView
	toggleItem := fyne.NewMenuItem(M("MainMenu.View.Toggle"), performToggle)
	toggleItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyF5,
		Modifier: fyne.KeyModifierShortcutDefault,
	}
	w.Canvas().AddShortcut(toggleItem.Shortcut, func(shortcut fyne.Shortcut) {
		performToggle()
	})
	dw.toggleLeftItem = fyne.NewMenuItem(M("MainMenu.View.ToggleLeft"), dw.commandToggleToolPanel)
	dw.toggleLeftItem.Checked = dw.toggleLeft
	dw.toggleRightItem = fyne.NewMenuItem(M("MainMenu.View.ToggleRight"), dw.commandToggleInspectorPanel)
	dw.toggleRightItem.Checked = dw.toggleRight

	view := fyne.NewMenu(M("MainMenu.View.Title"),
		toggleItem, dw.toggleLeftItem, dw.toggleRightItem,
	)
	dw.toggleItem = toggleItem
	dw.toggleRefresh = view.Refresh

	help := fyne.NewMenu(M("MainMenu.Help.Title"),
		fyne.NewMenuItem(M("MainMenu.Help.Designer.Site"), func() {
			goapp_fyne.OpenURL(DESIGNER_SITE)
		}),
		fyne.NewMenuItem(M("MainMenu.Help.Designer.Github"), func() {
			goapp_fyne.OpenURL(DESIGNER_GITHUB)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(M("MainMenu.Help.Fyne.Doc"), func() {
			goapp_fyne.OpenURL("https://developer.fyne.io")
		}),
		fyne.NewMenuItem(M("MainMenu.Help.Fyne.Support"), func() {
			goapp_fyne.OpenURL("https://fyne.io/support/")
		}),
	)

	main := fyne.NewMainMenu(
		wsp,
		file,
		view,
		help,
	)
	return main
}
