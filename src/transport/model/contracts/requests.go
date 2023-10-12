package contracts

// A structure representing data of a HTTP request for creating a new recognition app.
type CreateAppRequest struct {
	Name string `json:"app_name" validate:"required,min=3,max=200"`
}

// A structure representing data of a HTTP request for updating an existing recognition app.
type UpdateAppRequest struct {
	Name string `json:"new_app_name" validate:"required,min=3,max=200"`
}

// A structure representing data of a HTTP request for creating a new recognition workflow.
type CreateWorkflowRequest struct {
	Name                  string `json:"workflow_name" validate:"required,min=3,max=200"`
	AppId                 string `json:"app_id" validate:"uuid,required"`
	IsFullPageRecognition string `json:"is_full_page_recog" validate:"boolean,required"`
	SkipImageEnhancement  string `json:"skip_img_enhancement" validate:"boolean,required"`
	ExpectDifferentImages string `json:"expect_diff_images" validate:"boolean,required"`
}

// A structure representing data of a HTTP request for updating a specific workflow.
type UpdateWorkflowRequest struct {
	Name                  string `json:"new_workflow_name" validate:"required,min=3,max=200"`
	IsFullPageRecognition string `json:"is_full_page_recog" validate:"boolean,required"`
	SkipImageEnhancement  string `json:"skip_img_enhancement" validate:"boolean,required"`
	ExpectDifferentImages string `json:"expect_diff_images" validate:"boolean,required"`
}

type CreateTemplateRequest struct {
	Name   string               `json:"template_name" validate:"required,min=3,max=200"`
	Width  float32              `json:"width" validate:"required,number"`
	Height float32              `json:"height" validate:"required,number"`
	Fields []CreateFieldRequest `json:"fields" validate:"required,dive,required"`
}

type CreateFieldRequest struct {
	Name          string  `json:"field_name" validate:"required,min=3,max=200"`
	Width         float32 `json:"width" validate:"required,number"`
	Height        float32 `json:"height" validate:"required,number"`
	XPosition     float32 `json:"x_pos" validate:"required,number"`
	YPosition     float32 `json:"y_pos" validate:"required,number"`
	ExpectedValue string  `json:"expected_value" validate:"required,max=255"`
	IsIdentifying bool    `json:"is_identifying" validate:"required,boolean"`
}

// A structure representing data for a request for deleting users.
type DeleteUsersRequest struct {
	UserId string `json:"user_id" validate:"required,min=3"`
}

// A structure representing cloud events from MQ.
type CloudEvent[T interface{}] struct {
	Id     string `json:"id"`
	Data   T      `json:"data"`
	Source string `json:"source"`
}

// A structure representing data of a HTTP request for creating a new task group.
type CreateTaskGroupRequest struct {
	Name       string `json:"group_name" validate:"required,min=3,max=255"`
	WorkflowId string `json:"workflow_id" validate:"uuid,required"`
}

// A structure representing data of a HTTP request for obtaining a list of task groups for a specific workflow.
type GetTaskGroupsRequest struct {
	WorkflowId string `json:"workflow_id" validate:"uuid,required"`
}

// A structure representing data of a HTTP request for patch an existing task group.
type PatchTaskGroupRequest struct {
	Name string `json:"new_name" validate:"required,min=3,max=255"`
}

// A structure representing data of a HTTP request for creating a new task.
type CreateTaskRequest struct {
	Content     string `json:"task_content" validate:"required,min=3,max=600"`
	Description string `json:"description" validate:"required,min=3,max=255"`
	GroupId     string `json:"group_id" validate:"required,uuid"`
}
