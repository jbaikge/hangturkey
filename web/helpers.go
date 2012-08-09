package web

import (
	"html/template"
	"path/filepath"
	"strings"
)

var funcs = template.FuncMap{
	"trim": strings.TrimSpace,
}

func parseTemplates(files ...string) (t *template.Template) {
	name := filepath.Base(files[0])
	t = template.Must(template.New(name).Funcs(funcs).ParseFiles(files...))
	return
}
