package user

type ExistUserError struct{}

func (m *ExistUserError) Error() string {
	return "user already exists"
}
