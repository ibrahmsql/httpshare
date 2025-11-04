package template

import (
	"bytes"
	_ "embed"
	"html/template"

	"github.com/isa0-gh/httpshare/models"
)

//go:embed index.html
var indexHTML string

func Render(data models.DirectoryEntries) (string, error) {
	tmpl := template.Must(template.New("index").Parse(indexHTML))
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
