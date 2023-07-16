package errors

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrAlreadyExists struct {
	Message string
}

func (e ErrAlreadyExists) Error() string {
	return e.Message
}
