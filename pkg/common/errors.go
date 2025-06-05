package common

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
