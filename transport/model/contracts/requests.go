package contracts

// A structure representing data of a HTTP request for creating a new recognition app.
type CreateAppRequestData struct {
	Name string `json:"app_name" validate:"required, min=3, max=200"`
}
