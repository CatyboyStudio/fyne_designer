package main

import (
	"fyne_designer/widgets"
	"goapp_fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func testui() {
	var testobj *widgets.DesignCellWidget
	// l1 := widgets.NewDesignCellWidget(1, true, widget.NewLabel("Test"))
	// l2 := widgets.NewDesignCellWidget(2, true, widget.NewLabel("Test Hahaha\n Hahaha"))
	// c1 := widgets.NewDesignCellWidget(3, true, container.NewWithoutLayout(l1, l2))
	testobj = widgets.NewDesignCellWidget(1, widget.NewLabel("test"), func() {
		testobj.Active = !testobj.Active
		testobj.Refresh()
	})
	co2 := container.NewMax()
	co2.Add(testobj)

	co := container.NewMax()
	co.Add(container.NewVBox(co2))

	win := fyne.CurrentApp().NewWindow("Test")
	win.Resize(fyne.NewSize(740, 480))
	win.SetContent(co)
	goapp_fyne.ShowMaximizeWindow(win)
}
