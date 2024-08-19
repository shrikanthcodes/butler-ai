package templates

import (
	model "backend/internal/model"
	"fmt"
)

func initializeTemplateService() *TemplateService {
	// Initialize the template service
	templateService, err := NewTemplateService("templates/*.tmpl")
	if err != nil {
		fmt.Printf("Error initializing template service: %v\n", err)
		return nil
	}
	return templateService
}

// RenderRecipeTemplate renders the recipe template with the complete user data.
func RenderRecipeTemplate(data model.CompleteUserData) (string, error) {
	ts := initializeTemplateService()
	result, err := ts.RenderTemplate("recipe", data)
	if err != nil {
		return "", fmt.Errorf("error rendering recipe template: %w", err)
	}
	return result, nil
}
