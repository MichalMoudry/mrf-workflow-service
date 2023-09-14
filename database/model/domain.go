package model

import (
	"time"

	"github.com/google/uuid"
)

// A structure representing a recognition application.
type Application struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	CreatorId   string    `db:"creator_id"`
	DateAdded   time.Time `db:"date_added"`
	DateUpdated time.Time `db:"date_updated"`
}

// A constructor function for the Application structure.
func NewApplication(name, creatorId string) *Application {
	now := time.Now()
	return &Application{
		Id:          uuid.New(),
		Name:        name,
		CreatorId:   creatorId,
		DateAdded:   now,
		DateUpdated: now,
	}
}

// A structure representing a *..0 relationship between application and users.
type ApplicationUsers struct {
	Id            uuid.UUID
	ApplicationId uuid.UUID
	UserId        string
}

// A structure representing a recognition workflow/process.
type Workflow struct {
	Id            uuid.UUID
	Name          string
	ApplicationId uuid.UUID
	Templates     []DocumentTemplate
	Settings      WorkflowSetting
	DateAdded     time.Time
	DateUpdated   time.Time
}

// A constructor function for the Application structure.
func NewWorkflow(name string, appId uuid.UUID, settings WorkflowSetting) *Workflow {
	now := time.Now()
	return &Workflow{
		Id:          uuid.New(),
		Name:        name,
		Settings:    settings,
		DateAdded:   now,
		DateUpdated: now,
	}
}

// A structure representing a setting of a specific recognition workflow.
type WorkflowSetting struct {
	IsFullPageRecognition bool
	SkipImageEnhancement  bool
	ExpectDifferentImages bool
}

// A structure representing a template of a processed document.
type DocumentTemplate struct {
	Id          uuid.UUID
	Name        string
	Width       float32
	Height      float32
	Fields      []TemplateField
	DateAdded   time.Time
	DateUpdated time.Time
}

// A constructor function for the DocumentTemplate structure.
func NewDocumentTemplate(name string, width, height float32) *DocumentTemplate {
	now := time.Now()
	return &DocumentTemplate{
		Id:          uuid.New(),
		Name:        name,
		Width:       width,
		Height:      height,
		DateAdded:   now,
		DateUpdated: now,
	}
}

// A structure representing a field associated with a specific document template.
type TemplateField struct {
	Id            uuid.UUID
	Name          string
	Width         float32
	Height        float32
	XPosition     float32
	YPosition     float32
	ExpectedValue string
	IsIdentifying bool // A property for signaling if field is used during document identification.
	DateAdded     time.Time
	DateUpdated   time.Time
}

// A constructor function for the TemplateField structure.
func NewTemplateField(name string, width, height, xPos, yPos float32, expectedVal string, isId bool) *TemplateField {
	now := time.Now()
	return &TemplateField{
		Id:            uuid.New(),
		Name:          name,
		Width:         width,
		Height:        height,
		XPosition:     xPos,
		YPosition:     yPos,
		ExpectedValue: expectedVal,
		IsIdentifying: isId,
		DateAdded:     now,
		DateUpdated:   now,
	}
}
