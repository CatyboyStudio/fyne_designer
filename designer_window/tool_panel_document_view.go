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
func (th *DocumentViewItem) CreateRenderer() fyne.WidgetRenderer {
	th.label = widget.NewLabel("")
	th.build()
	buttons := container.NewHBox(
		canvas.NewRectangle(theme.ForegroundColor()),
		widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
			key := th.GetText()
			th.self.saveBy(key)
		}),
		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
			key := th.GetText()
			th.self.reloadBy(key)
		}),
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
			key := th.GetText()
			th.self.deleteBy(key)
		}),
	)
	co := container.NewBorder(nil, nil, nil, buttons, th.label)
	return widget.NewSimpleRenderer(co)
}

func (th *DocumentViewItem) GetText() string {
	text := ""
	if th.Data != nil {
		text, _ = th.Data.Get()
	}
	return text
}

func (th *DocumentViewItem) UpdateData(data binding.String) {
	th.Data = data
	th.build()
	th.label.Refresh()
}

func (th *DocumentViewItem) build() {
	if th.label != nil {
		s := th.GetText()
		if s != "" {
			_, s = DocumentViewItem_SplitData(s)
		}
		th.label.Text = s
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

func (th *DocumentView) Build() fyne.CanvasObject {
	list := widget.NewListWithData(
		th.docs,
		func() fyne.CanvasObject {
			return NewDocumentViewItem(th)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v := di.(binding.String)
			o := co.(*DocumentViewItem)
			o.UpdateData(v)
		},
	)
	th.docList = list
	list.OnSelected = func(id widget.ListItemID) {
		s, _ := th.docs.GetValue(id)
		docid, _ := DocumentViewItem_SplitData(s)
		InvokeWorkspaceTask(func(w *workspace.Workspace) (any, error) {
			return nil, w.ActiveDocument(docid, true)
		})
	}
	return list
}

func (th *DocumentView) addDocument(doc *workspace.Document) {
	ss := DocumentViewItem_MakeData(doc.GetId(), doc.GetTitle())
	if ss != "" {
		th.docs.Append(ss)
	}
}

func (th *DocumentView) removeDocument(doc *workspace.Document) {
	ss := DocumentViewItem_MakeData(doc.GetId(), doc.GetTitle())
	if ss != "" {
		slist, _ := th.docs.Get()
		nlist := arrutil.StringsRemove(slist, ss)
		th.docs.Set(nlist)
		th.docList.Refresh()
	}
}

func (th *DocumentView) deleteBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	InvokeWorkspaceTask(func(w *workspace.Workspace) (any, error) {
		return nil, w.CloseDocument(docid)
	})
}

func (th *DocumentView) saveBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	InvokeWorkspaceTask(func(w *workspace.Workspace) (any, error) {
		return nil, w.SaveDocument(docid)
	})
}

func (th *DocumentView) reloadBy(s string) {
	docid, _ := DocumentViewItem_SplitData(s)
	InvokeWorkspaceTask(func(w *workspace.Workspace) (any, error) {
		return nil, w.ReloadDocument(docid)
	})
}
