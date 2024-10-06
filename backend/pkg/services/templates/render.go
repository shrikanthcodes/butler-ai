package templates

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

// TemplateService is responsible for managing and rendering templates.
type TemplateService struct {
	templates *template.Template
}

// NewTemplateService initializes all templates from the provided directory and returns a TemplateService instance.
func NewTemplateService(directory string) (*TemplateService, error) {
	// Use filepath.Glob to get all .tmpl files in the directory
	pattern := filepath.Join(directory, "*.tmpl")
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
		return "", fmt.Errorf("failed to render template '%s': %w", templateName, err)
	}
	return buffer.String(), nil
}

const (
	RECIPE_MODE          = "recipe-mode-001.tmpl"
	SHOPPING_MODE        = "shopping-mode-001.tmpl"
	HEALTH_MODE          = "health-mode-001.tmpl"
	MOTIVATION_MODE      = "motivation-mode-001.tmpl"
	CALORIE_TRACKER_MODE = "calorie-tracker-mode-001.tmpl"
	SUMMARIZATION_MODE   = "summarization-mode-001.tmpl"
)
