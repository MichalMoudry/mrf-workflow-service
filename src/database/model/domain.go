package model

import (
	"time"

	"github.com/google/uuid"
)

// A structure representing a recognition application.
type Application struct {
	Id               uuid.UUID `db:"id"`
	Name             string    `db:"app_name"`
	CreatorId        string    `db:"creator_id"`
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
	DateAdded        time.Time `db:"date_added"`
	DateUpdated      time.Time `db:"date_updated"`
}

// A constructor function for the Application structure.
func NewApplication(name, creatorId string) *Application {
	now := time.Now()
	return &Application{
		Id:               uuid.New(),
		Name:             name,
		CreatorId:        creatorId,
		ConcurrencyStamp: uuid.New(),
		DateAdded:        now,
		DateUpdated:      now,
	}
}

type UserRole string

// A structure representing a *..0 relationship between application and users.
type ApplicationUsers struct {
	Id            uuid.UUID
	ApplicationId uuid.UUID
	UserId        string
	Role          UserRole
}

// A structure representing a recognition workflow/process.
type Workflow struct {
	Id                    uuid.UUID `db:"id"`
	Name                  string    `db:"workflow_name"`
	ApplicationId         uuid.UUID `db:"application_id"`
	Templates             []DocumentTemplate
	IsFullPageRecognition bool      `db:"setting_is_full_page_recog"`
	SkipImageEnhancement  bool      `db:"setting_skip_enhancement"`
	ExpectDifferentImages bool      `db:"setting_expect_diff_images"`
	ConcurrencyStamp      uuid.UUID `db:"concurrency_stamp"`
	DateAdded             time.Time `db:"date_added"`
	DateUpdated           time.Time `db:"date_updated"`
}

// A constructor function for the Application structure.
func NewWorkflow(name string, appId uuid.UUID, settings WorkflowSetting) *Workflow {
	now := time.Now()
	return &Workflow{
		Id:                    uuid.New(),
		Name:                  name,
		ApplicationId:         appId,
		IsFullPageRecognition: settings.IsFullPageRecognition,
		SkipImageEnhancement:  settings.SkipImageEnhancement,
		ExpectDifferentImages: settings.ExpectDifferentImages,
		ConcurrencyStamp:      uuid.New(),
		DateAdded:             now,
		DateUpdated:           now,
	}
}

// A structure representing a template of a processed document.
type DocumentTemplate struct {
	Id               uuid.UUID `db:"id"`
	Name             string    `db:"template_name"`
	Width            float32   `db:"width"`
	Height           float32   `db:"height"`
	Image            []byte    `db:"image"`
	Fields           []TemplateField
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
	DateAdded        time.Time `db:"date_added"`
	DateUpdated      time.Time `db:"date_updated"`
}

// A constructor function for the DocumentTemplate structure.
func NewDocumentTemplate(name string, width, height float32) *DocumentTemplate {
	now := time.Now()
	return &DocumentTemplate{
		Id:               uuid.New(),
		Name:             name,
		Width:            width,
		Height:           height,
		ConcurrencyStamp: uuid.New(),
		DateAdded:        now,
		DateUpdated:      now,
	}
}

// A structure representing a field associated with a specific document template.
type TemplateField struct {
	Id               uuid.UUID `db:"id"`
	Name             string    `db:"field_name"`
	Width            float32   `db:"width"`
	Height           float32   `db:"height"`
	XPosition        float32   `db:"x_position"`
	YPosition        float32   `db:"y_position"`
	ExpectedValue    string    `db:"expected_value"`
	IsIdentifying    bool      `db:"is_identifying"` // A property for signaling if field is used during document identification.
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
	DateAdded        time.Time `db:"date_added"`
	DateUpdated      time.Time `db:"date_updated"`
}

// A constructor function for the TemplateField structure.
func NewTemplateField(name string, width, height, xPos, yPos float32, expectedVal string, isId bool) *TemplateField {
	now := time.Now()
	return &TemplateField{
		Id:               uuid.New(),
		Name:             name,
		Width:            width,
		Height:           height,
		XPosition:        xPos,
		YPosition:        yPos,
		ExpectedValue:    expectedVal,
		IsIdentifying:    isId,
		ConcurrencyStamp: uuid.New(),
		DateAdded:        now,
		DateUpdated:      now,
	}
}