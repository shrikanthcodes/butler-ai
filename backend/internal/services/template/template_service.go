package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

// TemplateService is a struct that will handle templating logic.
type TemplateService struct {
	templates *template.Template
}

// NewTemplateService initializes all templates and returns a TemplateService instance.
func NewTemplateService(pattern string) (*TemplateService, error) {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}
	return &TemplateService{
		templates: tmpl,
	}, nil
}

// RenderTemplate renders a template with the given name and data.
func (ts *TemplateService) RenderTemplate(templateName string, data interface{}) (string, error) {
	var buffer bytes.Buffer
	err := ts.templates.ExecuteTemplate(&buffer, templateName, data)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}
	return buffer.String(), nil
}
