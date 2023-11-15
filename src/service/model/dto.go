package model

import (
	"time"

	"github.com/google/uuid"
)

type TaskGroupData struct {
	Id        uuid.UUID
	Name      string
	DateAdded time.Time
}

type TaskData struct {
	Id          uuid.UUID
	Name        string
	Description string
	Content     []byte
	DateAdded   time.Time
}

type WorkflowData struct {
	Id                    uuid.UUID `json:"workflow_id"`
	IsFullPageRecognition bool      `json:"is_full_page_recognition"`
	SkipImageEnhancement  bool      `json:"skip_img_enchancement"`
	ExpectDifferentImages bool      `json:"expect_diff_images"`
}
