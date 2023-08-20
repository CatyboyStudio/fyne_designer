package workspace

import (
	"path/filepath"
	"strings"
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

func (this *Workspace) OpenDocument(doc *Document) error {
	if o, ok := this.documents[doc.id]; ok {
		if o == doc {
			return nil
		}
		err := this.CloseDocument(o)
		if err != nil {
			return err
		}
	}

	// open document
	this.documents[doc.id] = doc
	if doc.Filepath != "" {
		n := filepath.Base(doc.Filepath)
		if strings.HasSuffix(n, DOC_EXT) {
			n = n[:len(n)-len(DOC_EXT)]
		}
		doc.title = n
	}
	this.RaiseEvent(EVENT_DOC_OPEN, doc)

	return nil
}

func (this *Workspace) CloseDocument(doc *Document) error {
	if old, ok := this.documents[doc.id]; ok {
		if old == doc {
			delete(this.documents, doc.id)
			doc.dispose()
			this.RaiseEvent(EVENT_DOC_CLOSE, doc)
			return nil
		}
	}
	return nil
}
