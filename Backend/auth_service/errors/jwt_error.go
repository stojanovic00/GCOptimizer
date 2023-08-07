package errors

type ErrInvalidToken struct{}

func (e ErrInvalidToken) Error() string {
	return "Token is invalid"
}
