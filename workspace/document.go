package workspace

import (
	"fmt"
	"goapp_commons"
	"noc"
	"path/filepath"
	"strings"

	. "cbsutil/valconv"
)

const DOC_EXT = ".doc.toml"
const DOC_COMTYPE = "document"

func init() {
	noc.RegisterGlobalFactory(DOC_COMTYPE, func(comtype string) (noc.Component, error) {
		return newDocument(), nil
	})
}

var _ (noc.Component) = (*Document)(nil)

type Document struct {
	noc.BaseComponent

	Filepath    string
	GenFilepath string
	PackageName string
	Root        *Element

	title string
}

func newDocument() *Document {
	return &Document{}
}

func (this *Document) OnCreate(info *noc.ComponentInfo) {
	this.BaseComponent.OnCreate(info)
	info.Flag.Set(noc.FLAG_DONT_DELETE)
	info.GetObject().MainData = this
}

func (this *Document) GetId() string {
	return this.Info().GetObject().Id()
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

func (this *Document) ToJson() (map[string]any, error) {
	ret := make(map[string]any)
	ret["package"] = this.PackageName
	ret["gen_file"] = this.GenFilepath
	return ret, nil
}

func (this *Document) FromJson(data map[string]any) error {
	jd := MapStringAny(data)
	this.PackageName = jd.Get("package").ToString().Value
	this.GenFilepath = jd.Get("gen_file").ToString().Value
	return nil
}
