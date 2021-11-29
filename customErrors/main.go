package customErrors

import "fmt"

type NotFoundError struct {
	Resource string
}

func NewNotFoundError(resource string) error {
	return &NotFoundError{Resource: resource}
}

func (nfe *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", nfe.Resource)
}

type ForbiddenUser struct {
	message string
}

func NewForbiddenUser(message string) error {
	return &ForbiddenUser{message: message}
}

func (fu *ForbiddenUser) Error() string {
	return fmt.Sprintf(fu.message)
}
