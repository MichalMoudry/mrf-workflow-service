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

// A structure representing a group of custom tasks.
type TaskGroup struct {
	Id          uuid.UUID
	Name        string
	Tasks       []Task
	DateAdded   time.Time
	DateUpdated time.Time
}

// A constructor function for the TaskGroup structure.
func NewTaskGroup(name string) *TaskGroup {
	now := time.Now()
	return &TaskGroup{
		Id:          uuid.New(),
		Name:        name,
		DateAdded:   now,
		DateUpdated: now,
	}
}

// A structure representing a custom processing task.
type Task struct {
	Id          uuid.UUID
	Name        string
	Content     []byte
	Description string
	DateAdded   time.Time
	DateUpdated time.Time
}

// A constructor function for the Task structure.
func NewTask(name, description string, content []byte) *Task {
	now := time.Now()
	return &Task{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Content:     content,
		DateAdded:   now,
		DateUpdated: now,
	}
}
