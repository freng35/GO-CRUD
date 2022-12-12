package order

type NoUserError struct{}

func (m *NoUserError) Error() string {
	return "user doesnt exist"
}

type NoBookError struct{}

func (m *NoBookError) Error() string {
	return "book doesnt exist"
}

type ZeroBooksError struct{}

func (m *ZeroBooksError) Error() string {
	return "zero book"
}
