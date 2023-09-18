package repositories

import (
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
)

type TemplateRepository struct{}

// A method for adding a new document template to the database.
func (TemplateRepository) AddTemplate(template *model.DocumentTemplate) (uuid.UUID, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return uuid.Nil, err
	}

	rows, err := ctx.NamedQuery(query.CreateTemplate, template)
	if err != nil {
		return uuid.Nil, err
	}
	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(returnedId)
}

// A method for deleting a specific template from the database.
func (TemplateRepository) DeleteTemplate(templateId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteTemplate, templateId); err != nil {
		return err
	}
	return nil
}

// Method for updating a specific template.
func (TemplateRepository) UpdateTemplate() error {
	return nil
}
