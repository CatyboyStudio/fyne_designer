package workspace

import (
	"bytes"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func (th *Workspace) SelectDocument(f func(doc *Document) bool) []*Document {
	var ret []*Document
	for _, doc := range th.documents {
		if f == nil || f(doc) {
			ret = append(ret, doc)
		}
	}
	return ret
}

func (th *Workspace) GetDocument(id string) *Document {
	return th.documents[id]
}

func (th *Workspace) OpenDocument(doc *Document) error {
	id := doc.GetId()
	if o, ok := th.documents[id]; ok {
		if o == doc {
			return nil
		}
		err := th.CloseDocument(id)
		if err != nil {
			return err
		}
	}

	// open document
	th.documents[id] = doc
	doc.updateTitle()
	th.RaiseEvent(EVENT_DOC_OPEN, doc)

	if th.activeDocument == nil {
		th.ActiveDocument(id, true)
	}

	return nil
}

func (th *Workspace) CloseDocument(id string) error {
	doc, ok := th.documents[id]
	if !ok {
		return nil
	}

	if th.activeDocument == doc {
		th.ActiveDocument(id, false)
	}

	delete(th.documents, id)
	th.RaiseEvent(EVENT_DOC_CLOSE, doc)
	th.node.DeleteObject(doc.Info().GetObject())
	return nil
}

func (th *Workspace) ReloadDocument(id string) error {
	time.Sleep(time.Second * 5)
	return nil
}

func (th *Workspace) ActiveDocument(id string, a bool) error {
	doc, ok := th.documents[id]
	if !ok {
		return nil
	}
	if a {
		if th.activeDocument != doc {
			th.activeDocument = doc
			th.RaiseEvent(EVENT_DOC_ACTIVE, doc)
		}
	} else {
		if th.activeDocument == doc {
			th.activeDocument = nil
			th.RaiseEvent(EVENT_DOC_ACTIVE, nil)
		}
	}
	return nil
}

func (th *Workspace) SaveDocument(id string) error {
	// TODO: 先保存Document，再保存其他对象
	doc, ok := th.documents[id]
	if !ok {
		return nil
	}
	if doc.Filepath == "" {
		th.NextEvent(EVENT_DOC_SAVEFILE, doc, func(w *Workspace) (any, error) {
			return nil, w.SaveDocument(id)
		})
		return nil
	}
	data, err := doc.ToJson()
	if err != nil {
		return errors.WithStack(err)
	}
	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(data)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.WriteFile(doc.Filepath, buf.Bytes(), 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (th *Workspace) LoadDocument(filename string) error {
	// TODO: 先加载Document，再加载其他对象
	bs, err := os.ReadFile(filename)
	if err != nil {
		return errors.WithStack(err)
	}
	data := make(map[string]any)
	err = toml.Unmarshal(bs, &data)
	if err != nil {
		return errors.WithStack(err)
	}
	obj, err := th.node.LoadObject(data)
	if err != nil {
		return errors.WithStack(err)
	}
	com, err := obj.AddComponent(DOC_COMTYPE)
	if err != nil {
		th.node.DeleteObject(obj)
		return errors.WithStack(err)
	}
	doc := com.(*Document)
	doc.Filepath = filename
	err = doc.FromJson(data)
	if err != nil {
		return err
	}
	return th.OpenDocument(doc)
}
