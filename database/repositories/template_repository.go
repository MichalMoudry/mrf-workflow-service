package repositories

import (
	"workflow-service/database/model"

	"github.com/google/uuid"
)

type TemplateRepository struct{}

// A method for adding a new document template to the database.
func (TemplateRepository) AddTemplate(template *model.DocumentTemplate) (uuid.UUID, error) {
	return template.Id, nil
}

// A method for deleting a specific template from the database.
func (TemplateRepository) DeleteTemplate(templateId uuid.UUID) error {
	return nil
}

// Method for adding a new field to an existing document template.
func AddField(field *model.TemplateField) (uuid.UUID, error) {
	return field.Id, nil
}

// A method for deleting a field of a specific document template.
func (TemplateRepository) DeleteField(fieldId uuid.UUID) error {
	return nil
}

// Method for updating an existing document template field.
func (TemplateRepository) UpdateField() error {
	return nil
}
