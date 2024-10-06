package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

// TemplateService is responsible for managing and rendering templates.
type TemplateService struct {
	templates *template.Template
}

// NewTemplateService initializes all templates from the provided pattern and returns a TemplateService instance.
func NewTemplateService(pattern string) (*TemplateService, error) {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}
	return &TemplateService{
		templates: tmpl,
	}, nil
}

// RenderTemplate renders the specified template using the provided data context.
func (ts *TemplateService) RenderTemplate(templateName string, data interface{}) (string, error) {
	var buffer bytes.Buffer
	if err := ts.templates.ExecuteTemplate(&buffer, templateName, data); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}
	return buffer.String(), nil
}
