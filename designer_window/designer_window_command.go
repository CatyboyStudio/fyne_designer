package designer_window

import (
	"fyne_designer/workspace"
	"goapp_commons"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (dw *DesignerWindow) syncToggleValue() {
	dw.toggleLeftItem.Checked = dw.toggleLeft
	dw.toggleRightItem.Checked = dw.toggleRight
	dw.toggleItem.Checked = dw.toggle
	dw.toggleRefresh()
}

func (dw *DesignerWindow) toggleValue(p *bool) {
	*p = !*p
	dw.syncToggleValue()
	dw.build_Designer()
}

func (dw *DesignerWindow) commandToggleView() {
	dw.toggleValue(&dw.toggle)
}

func (dw *DesignerWindow) commandToggleToolPanel() {
	dw.toggleValue(&dw.toggleLeft)
}

func (dw *DesignerWindow) commandToggleInspectorPanel() {
	dw.toggleValue(&dw.toggleRight)
}

func (dw *DesignerWindow) commandChangeDir() {
	workspace.PostWorkspaceTask(func(w *workspace.Workspace) (any, error) {
		dir := w.Dir()
		go dw.doChangeDir(dir)
		return nil, nil
	}, nil)
}

func (dw *DesignerWindow) doChangeDir(dir string) {
	M := goapp_commons.GetMessage
	input := widget.NewEntry()
	input.Text = dir
	button := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
			if err != nil {
				return
			}
			if lu == nil {
				return
			}
			s := lu.String()
			if strings.HasPrefix(s, "file://") {
				d, _ := filepath.Abs(s[len("file://"):])
				input.Text = d
				input.Refresh()
			}
		}, dw.window)
	})
	co1 := container.NewBorder(nil, nil, nil, button, input)
	item1 := widget.NewFormItem(M("Dialog.ChangeWorkspaceDir.Label"), container.NewMax(co1))
	dlg := dialog.NewForm(M("Dialog.ChangeWorkspaceDir.Title"),
		M("ConfirmDialog.Confirm"), M("ConfirmDialog.Dismiss"),
		[]*widget.FormItem{item1}, func(v bool) {
			if v {
				s := strings.TrimSpace(input.Text)
				workspace.PostWorkspaceTask(func(w *workspace.Workspace) (any, error) {
					w.SetDir(s)
					return nil, nil
				}, nil)
			}
		}, dw.window,
	)
	dlg.Resize(fyne.NewSize(600, 200))
	dlg.Show()
}

func (dw *DesignerWindow) commandNewDocument() {
	InvokeWorkspaceTask(func(w *workspace.Workspace) (any, error) {
		obj := w.Node().NewObject()
		com, err := obj.AddComponent(workspace.DOC_COMTYPE)
		if err != nil {
			w.Node().DeleteObject(obj)
			return nil, err
		}
		return nil, w.OpenDocument(com.(*workspace.Document))
	})
}
