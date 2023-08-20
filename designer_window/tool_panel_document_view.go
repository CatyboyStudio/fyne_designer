package designer_window

import (
	"fmt"
	"fyne_designer/workspace"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/CatyboyStudio/fyne_widgets"
	"github.com/gookit/goutil/arrutil"
)

func DocumentViewItem_SplitData(s string) (string, string) {
	idx := strings.Index(s, ":")
	if idx >= 0 {
		id := s[:idx]
		title := s[idx+1:]
		return id, title
	}
	return "", s
}

func DocumentViewItem_MakeData(id, title string) string {
	return fmt.Sprintf("%s:%s", id, title)
}

type DocumentViewItemBuilder struct {
	self *DocumentView
	Data binding.String

	label *widget.Label
}

func NewDocumentViewItemBuilder(self *DocumentView) fyne_widgets.WidgetBuilder {
	o := &DocumentViewItemBuilder{
		self: self,
	}
	return o
}

// Build implements fyne_widgets.WidgetBuilder.
func (this *DocumentViewItemBuilder) Build() fyne.CanvasObject {
	this.label = widget.NewLabel(this.GetText())
	buttons := container.NewHBox(
		canvas.NewRectangle(theme.ForegroundColor()),
		widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), nil),
		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), nil),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
	)
	return container.NewBorder(nil, nil, nil, buttons, this.label)
}

func (this *DocumentViewItemBuilder) GetText() string {
	text := ""
	if this.Data != nil {
		text, _ = this.Data.Get()
	}
	return text
}

func (this *DocumentViewItemBuilder) UpdateData(data binding.String) {
	this.Data = data
	s := this.GetText()
	if s != "" {
		_, s = DocumentViewItem_SplitData(s)
	}
	this.label.Text = s
}

type DocumentView struct {
	codes    binding.StringList
	codeList *widget.List
}

func NewDocumentView() *DocumentView {
	return &DocumentView{
		codes: binding.NewStringList(),
	}
}

func (this *DocumentView) Build() fyne.CanvasObject {
	list := widget.NewListWithData(
		this.codes,
		func() fyne.CanvasObject {
			b := NewDocumentViewItemBuilder(this)
			return fyne_widgets.NewStatefuleWidget(b)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v := di.(binding.String)
			b := fyne_widgets.GetStatefuleWidgetBuilder(co)
			if b != nil {
				if o, ok := b.(*DocumentViewItemBuilder); ok {
					o.UpdateData(v)
					co.Refresh()
				}
			}
		},
	)
	this.codeList = list
	return list
}

func (this *DocumentView) addDocument(doc *workspace.Document) {
	ss := DocumentViewItem_MakeData(doc.GetId(), doc.GetTitle())
	if ss != "" {
		this.codes.Append(ss)
	}
}

func (this *DocumentView) deleteBy(s string) {
	slist, _ := this.codes.Get()
	nlist := arrutil.Remove[string](slist, s)
	this.codes.Set(nlist)
}

func (this *DocumentView) saveBy(s string) {
	slist, _ := this.codes.Get()
	for i, v := range slist {
		if v == s {
			if i > 0 {
				slist[i], slist[i-1] = slist[i-1], slist[i]
				nlist := arrutil.CloneSlice[string](slist)
				this.codes.Set(nlist)
				this.codeList.Refresh()
				return
			}
		}
	}
}

func (this *DocumentView) reloadBy(s string) {
	slist, _ := this.codes.Get()
	for i, v := range slist {
		if v == s {
			if i < len(slist)-1 {
				slist[i], slist[i+1] = slist[i+1], slist[i]
				nlist := arrutil.CloneSlice[string](slist)
				this.codes.Set(nlist)
				this.codeList.Refresh()
				return
			}
		}
	}
}
