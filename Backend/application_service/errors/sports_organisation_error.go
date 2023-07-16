package errors

type ErrEmailTaken struct{}

func (e ErrEmailTaken) Error() string {
	return "Given email is taken"
}
