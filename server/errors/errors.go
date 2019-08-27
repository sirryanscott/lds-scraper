package errors

type generalError struct {
	errorString string
}

func (e generalError) Error() string {
	return e.errorString
}

func (e generalError) New(text string) error {
	return generalError{text}
}
