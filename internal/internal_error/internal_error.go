package internal_error

type InternalError struct {
	Message string
	Err     string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func NotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewINternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server_error",
	}
}
