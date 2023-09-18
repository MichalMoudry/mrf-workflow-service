package service

import (
	"context"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/service/model/ioc"
	"workflow-service/transport/model/contracts"

	"github.com/google/uuid"
)

type TemplateService struct {
	TemplateRepository ioc.ITemplateRepository
	FieldRepository    ioc.IFieldRepository
}

// A constructor function for TemplateService structure.
func NewTemplateService(templateRepo ioc.ITemplateRepository, fieldRepo ioc.IFieldRepository) *TemplateService {
	return &TemplateService{
		TemplateRepository: templateRepo,
		FieldRepository:    fieldRepo,
	}
}

// A method for creating a new document template in the system.
func (srvc TemplateService) CreateTemplate(ctx context.Context, data contracts.CreateTemplateRequest) (templateId uuid.UUID, err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() { err = database.EndTransaction(tx, err) }()

	templateId, err = srvc.TemplateRepository.AddTemplate(
		model.NewDocumentTemplate(data.Name, data.Width, data.Height),
	)

	fields := make([]*model.TemplateField, len(data.Fields))
	for index, field := range data.Fields {
		fields[index] = model.NewTemplateField(
			field.Name,
			field.Width,
			field.Height,
			field.XPosition,
			field.YPosition,
			field.ExpectedValue,
			field.IsIdentifying,
		)
	}
	err = srvc.FieldRepository.AddFields(fields)
	return
}
