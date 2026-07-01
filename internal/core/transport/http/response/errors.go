package core_http_response

type ErrorResponse struct {
	Error   string `json:"error"        example:"full_error_text"`
	Message string `json:"message"      example:"error_message"`
}
