package repositories

import (
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
)

type FieldRepository struct {
}

// Method for adding document template fields in bulk.
func (FieldRepository) AddFields(fields []*model.TemplateField) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.CreateField, fields); err != nil {
		return err
	}
	return nil
}

// Method for bulk deleting document template fields.
func (FieldRepository) DeleteFields(fieldIds []uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.DeleteField, fieldIds); err != nil {
		return err
	}
	return nil
}
