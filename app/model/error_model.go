package model

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	InternalError     = "500"
	NotFoundError     = "404"
	UnAuthorizedError = "401"
)

func Error(msg string, status string) ErrorResponse {
	return ErrorResponse{Status: status, Message: msg}
}
