package models

import "fmt"

// ErrInvalidCardCode represents an error that tag with such name already exists.
type ErrInvalidCardCode struct {
	CardCode string
}

// IsErrInvalidCardCode checks if an error is an ErrInvalidCardCode.
func IsErrInvalidCardCode(err error) bool {
	_, ok := err.(ErrInvalidCardCode)
	return ok
}

func (err ErrInvalidCardCode) Error() string {
	return fmt.Sprintf("invalid card code [code: %s]", err.CardCode)
}

// ErrDuplicateCardCode represents an error that tag with such name already exists.
type ErrDuplicateCardCode struct {
	CardCode string
}

// IsErrDuplicateCardCode checks if an error is an ErrDuplicateCardCode.
func IsErrDuplicateCardCode(err error) bool {
	_, ok := err.(ErrDuplicateCardCode)
	return ok
}

func (err ErrDuplicateCardCode) Error() string {
	return fmt.Sprintf("duplicate card code [code: %s]", err.CardCode)
}

// ErrDeckNotFound will be thrown if id cannot be found
type ErrDeckNotFound struct {
	Id string
}

// Error returns the error message
func (err ErrDeckNotFound) Error() string {
	return fmt.Sprintf("deck not found [Id: %s]", err.Id)
}

// IsErrDeckNotFound checks if an error is a ErrDeckNotFound.
func IsErrDeckNotFound(err error) bool {
	_, ok := err.(ErrDeckNotFound)
	return ok
}

// ErrInvalidId represents an error that tag with such name already exists.
type ErrInvalidId struct {
	Id string
}

// IsErrInvalidId checks if an error is an ErrInvalidId.
func IsErrInvalidId(err error) bool {
	_, ok := err.(ErrInvalidId)
	return ok
}

func (err ErrInvalidId) Error() string {
	return fmt.Sprintf("invalid id format [id: %s]", err.Id)
}

// ErrDeckRemainingExceeded will be thrown if id cannot be found
type ErrDeckRemainingExceeded struct {
	Count     int
	Remaining int
}

// Error returns the error message
func (err ErrDeckRemainingExceeded) Error() string {
	return fmt.Sprintf("the requested number of cards (%d) exceeds the cards in the deck (%d)", err.Count, err.Remaining)
}

// IsErrDeckRemainingExceeded checks if an error is a ErrDeckRemainingExceeded.
func IsErrDeckRemainingExceeded(err error) bool {
	_, ok := err.(ErrDeckRemainingExceeded)
	return ok
}
