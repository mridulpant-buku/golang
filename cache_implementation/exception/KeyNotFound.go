package exception

import "fmt"

// Golang has support for errors in a really simple way. Go functions returns errors as a second return value. That is the standard way of implementing and using errors in Go.
type KeyNotFoundError struct {
	Key string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key '%s' not found", e.Key)
}
