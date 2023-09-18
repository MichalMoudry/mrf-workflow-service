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
	Name                  string `json:"workflow_id" validate:"required,min=3,max=200"`
	AppId                 string `json:"app_id" validate:"uuid,required"`
	IsFullPageRecognition bool   `json:"is_full_page_recog" validate:"boolean,required"`
	SkipImageEnhancement  bool   `json:"skip_img_enhancement" validate:"boolean,required"`
	ExpectDifferentImages bool   `json:"expect_diff_images" validate:"boolean,required"`
}
