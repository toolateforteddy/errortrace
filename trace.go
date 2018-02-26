package errortrace

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Wrap wraps an error with the file & line number that it came from.  If the
// input error is nil, this will do nothing and return nil.
func Wrap(err error) error {
	if err == nil {
		return err
	}
	_, file, line, _ := runtime.Caller(1)
	return tracedError{file, line, err}
}

// Errorf creates an error that captures the file & line that it was created at.
// The error content is defined by the format string and args like fmt.Errorf.
func Errorf(format string, args ...interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	return tracedError{file, line, fmt.Errorf(format, args...)}
}

type tracedError struct {
	file string
	line int
	err  error
}

func (t tracedError) Error() string {
	return fmt.Sprintf("[%s:%d] %s", filepath.Base(t.file), t.line, t.err.Error())
}
