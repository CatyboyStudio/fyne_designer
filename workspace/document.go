package workspace

import (
	"fmt"
	"goapp_commons"
	"path/filepath"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
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

func (this *Document) updateTitle() {
	if this.Filepath != "" {
		n := filepath.Base(this.Filepath)
		if strings.HasSuffix(n, DOC_EXT) {
			n = n[:len(n)-len(DOC_EXT)]
		}
		this.title = n
	}
}

func (this *Document) ToJson() map[string]any {
	ret := make(map[string]any)
	ret["Id"] = this.id
	ret["Package"] = this.PackageName
	ret["GenFile"] = this.GenFilepath
	return ret
}

func (this *Document) FromJson(data map[string]any) error {
	this.id = goapp_commons.GetValue[string](data, "Id", "")
	if this.id == "" {
		return errors.New("Document miss [Id]")
	}
	this.PackageName = goapp_commons.GetValue[string](data, "Package", "")
	this.GenFilepath = goapp_commons.GetValue[string](data, "GenFile", "")
	return nil
}
