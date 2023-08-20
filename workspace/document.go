package workspace

import (
	"fmt"

	"github.com/CatyboyStudio/goapp_commons"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const DOC_EXT = ".doc.toml"

type Document struct {
	id          string
	Filepath    string
	GenFilepath string
	PackageName string
	Root        *Element

	title string
}

func NewDocument() *Document {
	return &Document{
		id: gonanoid.Must(),
	}
}

func (this *Document) dispose() {

}

func (this *Document) GetId() string {
	return this.id
}

func (this *Document) GetTitle() string {
	if this.title == "" {
		t := goapp_commons.GetMessage("Document.Untitle")
		if this.PackageName != "" {
			return fmt.Sprintf("%s(%s)", t, this.PackageName)
		}
		return t
	}
	return this.title
}
