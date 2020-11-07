package inter

import "html/template"

type View interface {
	Template() string
}

type TemplateBuilder = func(template *template.Template) (*template.Template, error)
