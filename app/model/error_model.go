package model

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	InternalError = "InternalError"
)

func Error(msg string, status string) ErrorResponse {
	return ErrorResponse{Status: status, Message: msg}
}
