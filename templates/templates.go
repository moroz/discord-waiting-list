package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html.tmpl
var fs embed.FS

var IndexTemplate = template.Must(template.New("index.html.tmpl").ParseFS(fs, "index.html.tmpl"))
