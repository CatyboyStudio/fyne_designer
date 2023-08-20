package workspace

type Document struct {
	Root        *Element
	GenFileName string
	PackageName string
}

func NewDocument() *Document {
	return &Document{}
}
