package designer_window

import (
	"fyne_designer/widgets"
	"fyne_widget/inspector"
	"goapp_fyne"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (dw *DesignerWindow) build_Main() fyne.CanvasObject {
	co := container.NewMax(
		dw.build_Designer(),
		dw.build_PopupMessageView(),
	)
	return co
}

func (dw *DesignerWindow) on_PopupMessageData() {
	list, _ := dw.popupMessage.Data.Get()
	dw.messageBox.RemoveAll()
	for _, v := range list {
		msg := v.(*goapp_fyne.PopupMessage)
		tex := widget.NewRichTextWithText(msg.Message)
		tex.Wrapping = fyne.TextWrapWord
		if msg.IsError {
			seg := tex.Segments[0].(*widget.TextSegment)
			seg.Style.ColorName = theme.ColorNameError
		}
		dw.messageBox.Add(tex)
	}
	dw.messageBox.Refresh()
	if len(list) > 0 {
		dw.messageShower.Show()
	} else {
		dw.messageShower.Hide()
	}
}

func (dw *DesignerWindow) build_PopupMessageView() fyne.CanvasObject {
	dw.messageBox = container.NewVBox()
	mbar := container.NewPadded(dw.messageBox)
	g := container.NewMax(canvas.NewRectangle(color.RGBA{128, 128, 128, 64}), mbar)
	s := container.NewVScroll(g)
	m := NewMessageBox(fyne.NewPos(-1, -1), s)
	dw.messageShower = s
	dw.messageShower.Hide()

	dw.popupMessage.Data.AddListener(binding.NewDataListener(dw.on_PopupMessageData))
	return m
}

func (dw *DesignerWindow) build_Designer() fyne.CanvasObject {
	if dw.designerContainer == nil {
		dw.designerContainer = container.NewMax()
	}
	dw.designerContainer.RemoveAll()
	if dw.toggle {
		content := dw.build_Designer_Content()
		dw.designerContainer.Add(content)
	} else {
		content := container.NewBorder(
			dw.build_Designer_Toolbar(),
			nil, nil, nil,
			dw.build_Designer_Content(),
		)
		dw.designerContainer.Add(content)
	}
	return dw.designerContainer
}

func (dw *DesignerWindow) build_Designer_Toolbar() fyne.CanvasObject {
	if dw.toolbar == nil {
		toolbar := widget.NewToolbar(
			widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
				dw.commandNewDocument()
			}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.HelpIcon(), func() {
				err := goapp_fyne.OpenURL(DESIGNER_SITE)
				if err != nil {
					dw.log.Error("Open help url", "error", err)
				}
			}),
		)
		dw.toolbar = toolbar
	}
	return dw.toolbar
}

func (dw *DesignerWindow) build_Designer_Content() fyne.CanvasObject {
	if dw.toggle {
		return dw.build_Designer_View()
	} else {
		if dw.split3 == nil {
			split := widgets.NewSplit3(dw.build_Tool_Panel(),
				dw.build_Designer_View(),
				dw.build_Inspector_Panel(),
			)
			split.OffsetL = 0.25
			split.OffsetT = 0.75
			dw.split3 = split
		}
		content := dw.split3
		dw.split3.SetVisible(false, dw.toggleLeft)
		dw.split3.SetVisible(true, dw.toggleRight)
		return content
	}
}

func (dw *DesignerWindow) build_Tool_Panel() fyne.CanvasObject {
	if dw.toolp == nil {
		dw.toolp = newToolPanel()
	}
	return dw.toolp.build()
}

func (dw *DesignerWindow) build_Designer_View() fyne.CanvasObject {
	return canvas.NewRectangle(color.White)
}

func (dw *DesignerWindow) build_Inspector_Panel() fyne.CanvasObject {
	if dw.inspector == nil {
		dw.inspector = inspector.NewInspector()
	}
	return dw.inspector
}
