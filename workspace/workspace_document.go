package workspace

import (
	"bytes"
	"io/ioutil"

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
	if o, ok := this.documents[doc.id]; ok {
		if o == doc {
			return nil
		}
		err := this.CloseDocument(o.id)
		if err != nil {
			return err
		}
	}

	// open document
	this.documents[doc.id] = doc
	doc.updateTitle()
	this.RaiseEvent(EVENT_DOC_OPEN, doc)

	if this.activeDocument == nil {
		this.ActiveDocument(doc.id, true)
	}

	return nil
}

func (this *Workspace) CloseDocument(id string) error {
	doc, ok := this.documents[id]
	if !ok {
		return nil
	}
	old, ok := this.documents[doc.id]
	if !ok {
		return nil
	}
	if old != doc {
		return nil
	}

	if this.activeDocument == doc {
		this.ActiveDocument(doc.id, false)
	}

	delete(this.documents, doc.id)
	doc.dispose()
	this.RaiseEvent(EVENT_DOC_CLOSE, doc)
	return nil
}

func (this *Workspace) ReloadDocument(id string) error {
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
	doc, ok := this.documents[id]
	if !ok {
		return nil
	}
	if doc.Filepath == "" {
		this.NextEvent(EVENT_DOC_SAVEFILE, doc, func(w *Workspace) error {
			return w.SaveDocument(doc.id)
		})
		return nil
	}
	data := doc.ToJson()
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(data)
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
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithStack(err)
	}
	data := make(map[string]any)
	err = toml.Unmarshal(bs, &data)
	if err != nil {
		return errors.WithStack(err)
	}
	doc := NewDocument()
	doc.Filepath = filename
	err = doc.FromJson(data)
	if err != nil {
		return err
	}
	return this.OpenDocument(doc)
}
