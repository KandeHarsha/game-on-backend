package schema

import "errors"

type ErrorResponse struct {
	ErroCode    int    `json:"ErrorCode"`
	Message     string `json:"Message"`
	Description string `json:"Description"`
	ErrorInfo   string `json:"ErrorInfo"`
}

func GetSomethingWentWrongError() *ErrorResponse {
	return &ErrorResponse{
		ErroCode:    500,
		Message:     "Something went wrong",
		Description: "Please try again later",
		ErrorInfo:   "Something went wrong",
	}
}

func (e *ErrorResponse) Error() error {
	return errors.New(e.Description)
}
