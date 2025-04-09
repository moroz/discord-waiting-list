package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html.tmpl
var fs embed.FS

var IndexTemplate = template.Must(template.ParseFS(fs, "layout.html.tmpl", "index.html.tmpl"))
