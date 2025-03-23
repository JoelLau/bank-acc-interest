package async

import (
	"errors"
	"time"
)

var ErrTimeOut = errors.New("operation timed out")

func RunWithTimeOut(fn func() error, timeout time.Duration) error {
	done := make(chan error, 1) // Use an error channel instead of modifying err

	go func() {
		done <- fn()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return ErrTimeOut
	}
}
