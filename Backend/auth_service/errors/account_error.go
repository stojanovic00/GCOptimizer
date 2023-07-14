package errors

type ErrEmailTaken struct{}

func (e ErrEmailTaken) Error() string {
	return "Given email is taken"
}

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrBadCredentials struct{}

func (e ErrBadCredentials) Error() string {
	return "Email or password incorrect"
}
