package designer_window

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/CatyboyStudio/goapp_commons"
)

const DESIGNER_SITE = "https://catyboy.itch.io/"
const DESIGNER_GITHUB = "https://github.com/CatyboyStudio/FyneDesigner"

func (this *DesignerWindow) build_MainMenu() *fyne.MainMenu {
	a := fyne.CurrentApp()
	w := this.window
	M := goapp_commons.GetMessage

	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	otherItem := fyne.NewMenuItem("Other", nil)
	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)
	fileItem := fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") })
	fileItem.Icon = theme.FileIcon()
	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)

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

	quitItem := fyne.NewMenuItem(M("MainMenu.File.Quit"), func() {
		a.Quit()
	})
	quitItem.IsQuit = true
	file := fyne.NewMenu(M("MainMenu.File.Title"),
		newItem, checkedItem, disabledItem,
		// fyne.NewMenuItemSeparator(), settingsItem,
		fyne.NewMenuItemSeparator(), quitItem,
	)

	performToggle := this.toggleView
	toggleItem := fyne.NewMenuItem(M("MainMenu.View.Toggle"), performToggle)
	toggleItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyF5,
		Modifier: fyne.KeyModifierShortcutDefault,
	}
	w.Canvas().AddShortcut(toggleItem.Shortcut, func(shortcut fyne.Shortcut) {
		performToggle()
	})
	this.toggleLeftItem = fyne.NewMenuItem(M("MainMenu.View.ToggleLeft"), this.toggleLeftPanel)
	this.toggleLeftItem.Checked = this.toggleLeft
	this.toggleRightItem = fyne.NewMenuItem(M("MainMenu.View.ToggleRight"), this.toggleRightPanel)
	this.toggleRightItem.Checked = this.toggleRight

	view := fyne.NewMenu(M("MainMenu.View.Title"),
		toggleItem, this.toggleLeftItem, this.toggleRightItem,
	)
	this.toggleItem = toggleItem
	this.toggleRefresh = view.Refresh

	help := fyne.NewMenu(M("MainMenu.Help.Title"),
		fyne.NewMenuItem(M("MainMenu.Help.Designer.Site"), func() {
			fyne_widgets.OpenURL(DESIGNER_SITE)
		}),
		fyne.NewMenuItem(M("MainMenu.Help.Designer.Github"), func() {
			fyne_widgets.OpenURL(DESIGNER_GITHUB)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(M("MainMenu.Help.Fyne.Doc"), func() {
			fyne_widgets.OpenURL("https://developer.fyne.io")
		}),
		fyne.NewMenuItem(M("MainMenu.Help.Fyne.Support"), func() {
			fyne_widgets.OpenURL("https://fyne.io/support/")
		}),
	)

	main := fyne.NewMainMenu(
		file,
		view,
		help,
	)
	checkedItem.Action = func() {
		checkedItem.Checked = !checkedItem.Checked
		main.Refresh()
	}
	return main
}
