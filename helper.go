package main

import (
	"io"
	"strings"
	"text/template"
)

var helperTmpl = `package {{ .PackageName }}

import (
	"net/http"

	"github.com/gorilla/context"
)

type {{ .KeyType }} int

const {{ .KeyName }} {{ .KeyType }} = 0

func Get{{ .TitleName }}(r *http.Request) {{ .Type }} {
	if rv := context.Get(r, {{ .KeyName }}); rv != nil {
		return rv.({{ .Type }})
	}
	return nil
}

func Set{{ .TitleName }}(r *http.Request, val {{ .Type }}) {
	context.Set(r, {{ .KeyName }}, val)
}`

type Helper struct {
	Name        string
	PackageName string
}

func (h Helper) Render(out io.Writer) error {
	t, err := template.New("helper").Parse(helperTmpl)
	if err != nil {
		return err
	}

	return t.Execute(out, h)
}

func (h Helper) TitleName() string {
	return strings.Trim(strings.Title(h.Name), " *")
}

func (h Helper) LowerName() string {
	return strings.Trim(strings.ToLower(h.Name), " *")
}

func (h Helper) KeyType() string {
	return h.LowerName() + "HelperKey"
}

func (h Helper) KeyName() string {
	return h.LowerName() + "Key"
}

func (h Helper) Type() string {
	return "*" + h.Name
}

func (h Helper) FileName() string {
	return h.LowerName() + "_helpers.go"
}
