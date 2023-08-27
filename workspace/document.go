package workspace

import (
	"fmt"
	"goapp_commons"
	"noc"
	"path/filepath"
	"strings"

	V "cbsutil/valconv"
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

func (th *Document) OnCreate(info *noc.ComponentInfo) {
	th.BaseComponent.OnCreate(info)
	info.Flag.Set(noc.FLAG_DONT_DELETE)
	info.GetObject().MainData = th
}

func (th *Document) GetId() string {
	return th.Info().GetObject().Id()
}

func (th *Document) GetTitle() string {
	if th.title == "" {
		t := goapp_commons.GetMessage("Document.Untitle")
		if th.PackageName != "" {
			return fmt.Sprintf("%s(%s)", t, th.PackageName)
		}
		return t
	}
	return th.title
}

func (th *Document) updateTitle() {
	if th.Filepath != "" {
		n := filepath.Base(th.Filepath)
		n = strings.TrimSuffix(n, DOC_EXT)
		th.title = n
	}
}

func (th *Document) ToJson() (map[string]any, error) {
	ret := make(map[string]any)
	ret["package"] = th.PackageName
	ret["gen_file"] = th.GenFilepath
	return ret, nil
}

func (th *Document) FromJson(data map[string]any) error {
	jd := V.MapStringAny(data)
	th.PackageName = jd.Get("package").ToString().Value
	th.GenFilepath = jd.Get("gen_file").ToString().Value
	return nil
}
