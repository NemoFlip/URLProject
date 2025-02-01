package payload

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkResponse struct {
}
