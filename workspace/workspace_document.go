package workspace

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func (this *Workspace) SelectDocument(f func(doc *Document) bool) []*Document {
	var ret []*Document
	for _, doc := range this.documents {
		if f == nil || f(doc) {
			ret = append(ret, doc)
		}
	}
	return ret
}

func (this *Workspace) GetDocument(id string) *Document {
	return this.documents[id]
}

func (this *Workspace) OpenDocument(doc *Document) error {
	id := doc.GetId()
	if o, ok := this.documents[id]; ok {
		if o == doc {
			return nil
		}
		err := this.CloseDocument(id)
		if err != nil {
			return err
		}
	}

	// open document
	this.documents[id] = doc
	doc.updateTitle()
	this.RaiseEvent(EVENT_DOC_OPEN, doc)

	if this.activeDocument == nil {
		this.ActiveDocument(id, true)
	}

	return nil
}

func (this *Workspace) CloseDocument(id string) error {
	doc, ok := this.documents[id]
	if !ok {
		return nil
	}
	old, ok := this.documents[id]
	if !ok {
		return nil
	}
	if old != doc {
		return nil
	}

	if this.activeDocument == doc {
		this.ActiveDocument(id, false)
	}

	delete(this.documents, id)
	this.RaiseEvent(EVENT_DOC_CLOSE, doc)
	this.DeleteObject(doc.Info().GetObject())
	return nil
}

func (this *Workspace) ReloadDocument(id string) error {
	time.Sleep(time.Second * 5)
	return nil
}

func (this *Workspace) ActiveDocument(id string, a bool) error {
	doc, ok := this.documents[id]
	if !ok {
		return nil
	}
	if a {
		if this.activeDocument != doc {
			this.activeDocument = doc
			this.RaiseEvent(EVENT_DOC_ACTIVE, doc)
		}
	} else {
		if this.activeDocument == doc {
			this.activeDocument = nil
			this.RaiseEvent(EVENT_DOC_ACTIVE, nil)
		}
	}
	return nil
}

func (this *Workspace) SaveDocument(id string) error {
	// TODO: 先保存Document，再保存其他对象
	doc, ok := this.documents[id]
	if !ok {
		return nil
	}
	if doc.Filepath == "" {
		this.NextEvent(EVENT_DOC_SAVEFILE, doc, func(w *Workspace) error {
			return w.SaveDocument(id)
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
	err = ioutil.WriteFile(doc.Filepath, buf.Bytes(), 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (this *Workspace) LoadDocument(filename string) error {
	// TODO: 先加载Document，再加载其他对象
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithStack(err)
	}
	data := make(map[string]any)
	err = toml.Unmarshal(bs, &data)
	if err != nil {
		return errors.WithStack(err)
	}
	obj, err := this.LoadObject(data)
	if err != nil {
		return errors.WithStack(err)
	}
	com, err := obj.AddComponent(DOC_COMTYPE)
	if err != nil {
		this.DeleteObject(obj)
		return errors.WithStack(err)
	}
	doc := com.(*Document)
	doc.Filepath = filename
	err = doc.FromJson(data)
	if err != nil {
		return err
	}
	return this.OpenDocument(doc)
}
