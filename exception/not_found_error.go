package exception

type NotFoundError struct {
	NotFoundError string
	Model         string
}

func (e *NotFoundError) Error() string {
	return e.NotFoundError
}

func NewNotFoundError(error string, model string) *NotFoundError {
	return &NotFoundError{NotFoundError: error, Model: model}
}