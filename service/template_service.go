package service

import "workflow-service/service/model/ioc"

type TemplateService struct {
	TemplateRepository ioc.ITemplateRepository
}

// A constructor function for TemplateService structure.
func NewTemplateService(templateRepo ioc.ITemplateRepository) *TemplateService {
	return &TemplateService{
		TemplateRepository: templateRepo,
	}
}
