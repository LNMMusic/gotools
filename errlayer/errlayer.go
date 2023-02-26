package errlayer

import "fmt"

// ____________________________________________________________________
// ErrorLayer: struct for error layer that implements error interface
type ErrorLayer struct {
	message string
	err    error
}
func (e *ErrorLayer) Error() string {
	msg := e.message
	if e.err != nil {
		msg += ". " + e.err.Error()
	}

	return msg
}

func New(message string) error {
	return &ErrorLayer{message: message}
}


// ____________________________________________________________________
// Wrap: wrap error with message
func Wrap(err error, message string) error {
	return &ErrorLayer{message: message, err: err}
}
func Wrapf(err error, format string, args ...interface{}) error {
	return &ErrorLayer{message: fmt.Sprintf(format, args...), err: err}
}

// Unwrap: unwrap error to previous error
func Unwrap(err error) error {
	if err, ok := err.(*ErrorLayer); ok {
		return err.err
	}
	return nil
}

// Target: check if target error is in chain of errors
func Target(err error, target error) bool {
	// check target
	if target == nil {
		return false
	}

	// check chain of errors
	for err != nil {
		if err == target {
			return true
		}
		err = Unwrap(err)
	}

	return false
}