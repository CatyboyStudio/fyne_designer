package workspace

import (
	"fyne_widget/inspector"
	"goapp_commons"
	"path/filepath"
)

var _ (inspector.EditorBuilder) = (*Document)(nil)

// BuildEditor implements inspector.EditorBuilder.
func (th *Document) BuildEditor(ed *inspector.Editor) error {

	M := goapp_commons.GetMessage
	p1 := inspector.String[string](M("Document.Editor.File"), func() (string, error) {
		fn := th.Filepath
		if fn != "" {
			ws := th.GetWorkspace()
			nfn, err := filepath.Rel(ws.Dir(), fn)
			if err == nil {
				fn = nfn
			}
		}
		return fn, nil
	}, func(s string) error {
		return nil
	})
	inspector.NewTextItem(p1).Bind(ed).Watch()

	p2 := inspector.StringRef(M("Document.Editor.GenFilepath"), &th.GenFilepath).
		WithOnUpdate(func() {
			ws := th.GetWorkspace()
			ws.RaiseEvent(EVENT_DOC_UPDATE, th)
		})
	inspector.NewFilePathItem(p2).Bind(ed)

	p3 := inspector.StringRef(M("Document.Editor.PackageName"), &th.PackageName).
		WithOnUpdate(func() {
			ws := th.GetWorkspace()
			ws.RaiseEvent(EVENT_DOC_UPDATE, th)
		})
	inspector.NewEntryItem(p3).Bind(ed)

	p4 := inspector.StringRef(M("Document.Editor.FuncName"), &th.FuncName)
	inspector.NewEntryItem(p4).Bind(ed)

	return nil
}
