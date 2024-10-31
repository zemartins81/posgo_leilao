package internal_error

type InternalError struct {
	Message string
	Err     string
}

func NotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewINternalServerError() {
	
}
