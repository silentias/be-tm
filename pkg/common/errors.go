package common

import "fmt"

type AppError struct {
	Message    string
	StatusCode int
}

func InternalServerError() *AppError {
	return &AppError{
		Message:    "Internal server error",
		StatusCode: 500,
	}
}

func ExistsError(field string) *AppError {
	return &AppError{
		Message:    fmt.Sprintf("This %v already exists", field),
		StatusCode: 400,
	}
}
