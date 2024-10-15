package template

import (
	"bytes"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"path/filepath"
	"text/template"
)

// TemplateService is responsible for managing and rendering templates.
type TemplateService struct {
	templates *template.Template
	log       *logger.Logger
}

// NewTemplateService initializes all templates from the provided directory and returns a TemplateService instance.
func NewTemplateService(log *logger.Logger) (*TemplateService, error) {
	// Use filepath.Glob to get all .tmpl files in the directory
	pattern := filepath.Join("internal", "resources", "*.tmpl")
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}
	return &TemplateService{
		templates: tmpl,
		log:       log,
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
	RecipeMode            = "recipe-mode-001.tmpl"
	ShoppingMode          = "shopping-mode-001.tmpl"
	HealthMode            = "health-mode-001.tmpl"
	MotivationMode        = "motivation-mode-001.tmpl"
	CalorieTrackerMode    = "calorie-tracker-mode-001.tmpl"
	ChatSummarizationMode = "chat-summarization-mode-001.tmpl"
)
