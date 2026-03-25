package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Username string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("User '%s' not found", e.Username)
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}

func FindUser(username string) (string, error) {

	if username == "" {
		return "", &ValidationError{Field: "username", Message: "Username cannot be empty"}
	}
	if username != "alice" && username != "bob" {
		return "", &NotFoundError{Username: username}
	}
	return fmt.Sprintf("User '%s' found", username), nil
}

func main() {
	for _, v := range []string{"", "charlie", "alice", "bob"} {
		result, err := FindUser(v)
		var valErr *ValidationError
		var notFoundErr *NotFoundError
		if errors.As(err, &valErr) {
			fmt.Printf("Validation error: %s\n", valErr.Message)
		} else if errors.As(err, &notFoundErr) {
			fmt.Printf("Not found error: %s\n", notFoundErr.Username)
		} else if err != nil {
			fmt.Printf("Unexpected error: %s\n", err)
		} else {
			fmt.Println(result)
		}
	}
}
