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

type DocumentViewItem struct {
	widget.BaseWidget
	self *DocumentView
	Data binding.String

	label *widget.Label
}

func NewDocumentViewItem(self *DocumentView) *DocumentViewItem {
	o := &DocumentViewItem{
		self: self,
	}
	o.ExtendBaseWidget(o)
	return o
}

// CreateRenderer implements fyne.Widget.
func (this *DocumentViewItem) CreateRenderer() fyne.WidgetRenderer {
	this.label = widget.NewLabel("")
	this.build()
	buttons := container.NewHBox(
		canvas.NewRectangle(theme.ForegroundColor()),
		widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
			key := this.GetText()
			this.self.saveBy(key)
		}),
		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
			key := this.GetText()
			this.self.reloadBy(key)
		}),
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
			key := this.GetText()
			this.self.deleteBy(key)
		}),
	)
	co := container.NewBorder(nil, nil, nil, buttons, this.label)
	return widget.NewSimpleRenderer(co)
}

func (this *DocumentViewItem) GetText() string {
	text := ""
	if this.Data != nil {
		text, _ = this.Data.Get()
	}
	return text
}

func (this *DocumentViewItem) UpdateData(data binding.String) {
	this.Data = data
	this.build()
	this.label.Refresh()
}

func (this *DocumentViewItem) build() {
	if this.label != nil {
		s := this.GetText()
		if s != "" {
			_, s = DocumentViewItem_SplitData(s)
		}
		this.label.Text = s
	}
}

type DocumentView struct {
	docs    binding.StringList
	docList *widget.List
}

func NewDocumentView() *DocumentView {
	return &DocumentView{
		docs: binding.NewStringList(),
	}
}

func (this *DocumentView) Build() fyne.CanvasObject {
	list := widget.NewListWithData(
		this.docs,
		func() fyne.CanvasObject {
			return NewDocumentViewItem(this)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v := di.(binding.String)
			o := co.(*DocumentViewItem)
			o.UpdateData(v)
		},
	)
	this.docList = list
	list.OnSelected = func(id widget.ListItemID) {
		s, _ := this.docs.GetValue(id)
		docid, _ := DocumentViewItem_SplitData(s)
		ExecWorkspaceTask(func(w *workspace.Workspace) error {
			return w.ActiveDocument(docid, true)
		})
	}
	return list
}

func (this *DocumentView) addDocument(doc *workspace.Document) {
	ss := DocumentViewItem_MakeData(doc.GetId(), doc.GetTitle())
	if ss != "" {
		this.docs.Append(ss)
	}
}

func (this *DocumentView) removeDocument(doc *workspace.Document) {
	ss := DocumentViewItem_MakeData(doc.GetId(), doc.GetTitle())
	if ss != "" {
		slist, _ := this.docs.Get()
		nlist := arrutil.StringsRemove(slist, ss)
		this.docs.Set(nlist)
		this.docList.Refresh()
	}
}

func (this *DocumentView) deleteBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	ExecWorkspaceTask(func(w *workspace.Workspace) error {
		return w.CloseDocument(docid)
	})
}

func (this *DocumentView) saveBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	ExecWorkspaceTask(func(w *workspace.Workspace) error {
		return w.SaveDocument(docid)
	})
}

func (this *DocumentView) reloadBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	ExecWorkspaceTask(func(w *workspace.Workspace) error {
		return w.ReloadDocument(docid)
	})
}
