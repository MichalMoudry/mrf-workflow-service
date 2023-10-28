package model

import (
	"time"

	"github.com/google/uuid"
)

// A structure encapsulating available info of a recognition app.
type ApplicationInfo struct {
	Id               uuid.UUID `db:"id"`
	Creator          string    `db:"creator_id"`
	Name             string    `db:"app_name"`
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
	Added            time.Time `db:"date_added"`
	Updated          time.Time `db:"date_updated"`
}

// A structure encapsulating available info of a specific workflow.
type WorkflowInfo struct {
	Id                    uuid.UUID `db:"id"`
	Name                  string    `db:"workflow_name"`
	IsFullPageRecognition bool      `db:"setting_is_full_page_recog"`
	SkipImageEnhancement  bool      `db:"setting_skip_enhancement"`
	ExpectDifferentImages bool      `db:"setting_expect_diff_images"`
	ConcurrencyStamp      uuid.UUID `db:"concurrency_stamp"`
	Added                 time.Time `db:"date_added"`
	Updated               time.Time `db:"date_updated"`
}

// A structure representing a setting of a specific recognition workflow.
type WorkflowSetting struct {
	IsFullPageRecognition bool `json:"is_full_page_recognition"`
	SkipImageEnhancement  bool `json:"skip_img_enchancement"`
	ExpectDifferentImages bool `json:"expecte_diff_images"`
}
