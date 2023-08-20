package designer_window

import (
	"fyne_designer/widgets"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/fyne_widgets"
)

func (this *DesignerWindow) build_Main() fyne.CanvasObject {
	co := container.NewMax(
		this.build_Designer(),
		this.build_PopupMessageView(),
	)
	return co
}

func (this *DesignerWindow) on_PopupMessageData() {
	list, _ := this.popupMessage.Data.Get()
	this.messageBox.RemoveAll()
	for _, v := range list {
		msg := v.(*fyne_widgets.PopupMessage)
		tex := widget.NewRichTextWithText(msg.Message)
		tex.Wrapping = fyne.TextWrapWord
		if msg.IsError {
			seg := tex.Segments[0].(*widget.TextSegment)
			seg.Style.ColorName = theme.ColorNameError
		}
		this.messageBox.Add(tex)
	}
	this.messageBox.Refresh()
	if len(list) > 0 {
		this.messageShower.Show()
	} else {
		this.messageShower.Hide()
	}
}

func (this *DesignerWindow) build_PopupMessageView() fyne.CanvasObject {
	this.messageBox = container.NewVBox()
	mbar := container.NewPadded(this.messageBox)
	g := container.NewMax(canvas.NewRectangle(color.RGBA{128, 128, 128, 64}), mbar)
	s := container.NewVScroll(g)
	m := NewMessageBox(fyne.NewPos(-1, -1), s)
	this.messageShower = s
	this.messageShower.Hide()

	this.popupMessage.Data.AddListener(binding.NewDataListener(this.on_PopupMessageData))
	return m
}

func (this *DesignerWindow) build_Designer() fyne.CanvasObject {
	if this.designerContainer == nil {
		this.designerContainer = container.NewMax()
	}
	this.designerContainer.RemoveAll()
	if this.toggle {
		content := this.build_Designer_Content()
		this.designerContainer.Add(content)
	} else {
		content := container.NewBorder(
			this.build_Designer_Toolbar(),
			nil, nil, nil,
			this.build_Designer_Content(),
		)
		this.designerContainer.Add(content)
	}
	return this.designerContainer
}

func (this *DesignerWindow) build_Designer_Toolbar() fyne.CanvasObject {
	if this.toolbar == nil {
		toolbar := widget.NewToolbar(
			widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
				log.Println("New document")
			}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.HelpIcon(), func() {
				err := fyne_widgets.OpenURL(DESIGNER_SITE)
				if err != nil {
					this.log.Error().Err(err).Msg("Open help url")
				}
			}),
		)
		this.toolbar = toolbar
	}
	return this.toolbar
}

func (this *DesignerWindow) build_Designer_Content() fyne.CanvasObject {
	if this.toggle {
		return this.build_Designer_View()
	} else {
		if this.split3 == nil {
			split := widgets.NewSplit3(this.build_Tool_Panel(), this.build_Designer_View(), widget.NewLabel("Right"))
			split.OffsetL = 0.25
			split.OffsetT = 0.75
			this.split3 = split
		}
		content := this.split3
		this.split3.SetVisible(false, this.toggleLeft)
		this.split3.SetVisible(true, this.toggleRight)
		return content
	}
}

func (this *DesignerWindow) build_Tool_Panel() fyne.CanvasObject {
	if this.toolp == nil {
		this.toolp = newToolPanel()
	}
	return this.toolp.build()
}

func (this *DesignerWindow) build_Designer_View() fyne.CanvasObject {
	return canvas.NewRectangle(color.White)
}
