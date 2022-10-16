package sentinel

import (
	"fmt"
)

// withSentinel holds a sentinel error along with the original cause of the error, plus a message.
type withSentinel struct {
	original error
	sentinel error
	msg      string
}

func (w withSentinel) Error() string {
	return fmt.Sprintf("%v: %s: %v", w.sentinel, w.msg, w.original)
}

// Wrap wraps an original error with a given sentinel error.
func Wrap(original error, sentinel error) error {
	//return errors.Errorf("%w:%w", original, sentinel)
	return withSentinel{
		original: original,
		sentinel: sentinel,
	}
}

// Unwrap returns the sentinel error.
func (w withSentinel) Unwrap() error {
	return w.sentinel
}

// WithMessage wraps an error with a given sentinel, with some context via a message.
func WithMessage(original, sentinel error, message string) error {
	return withSentinel{
		original: original,
		sentinel: sentinel,
		msg:      message,
	}
}

// WithMessagef wraps an error with a given sentinel, followed by some formatted message.
func WithMessagef(original error, sentinel error, format string, args ...interface{}) error {
	return WithMessage(original, sentinel, fmt.Sprintf(format, args))
}
