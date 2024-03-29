package models

import "fmt"

// InvalidCardCodeError represents an error that card code is in invalid format.
type InvalidCardCodeError struct {
	CardCode string
}

// IsInvalidCardCodeError checks if an error is an InvalidCardCodeError.
func IsInvalidCardCodeError(err error) bool {
	_, ok := err.(InvalidCardCodeError)

	return ok
}

// Error returns the error message.
func (err InvalidCardCodeError) Error() string {
	return fmt.Sprintf("invalid card code [code: %s]", err.CardCode)
}

// DuplicateCardCodeError represents an error that card code exist twice in cards filter.
type DuplicateCardCodeError struct {
	CardCode string
}

// IsDuplicateCardCodeError checks if an error is an DuplicateCardCodeError.
func IsDuplicateCardCodeError(err error) bool {
	_, ok := err.(DuplicateCardCodeError)

	return ok
}

// Error returns the error message.
func (err DuplicateCardCodeError) Error() string {
	return fmt.Sprintf("duplicate card code [code: %s]", err.CardCode)
}

// DeckNotFoundError will be thrown if id cannot be found.
type DeckNotFoundError struct {
	ID string
}

// Error returns the error message.
func (err DeckNotFoundError) Error() string {
	return fmt.Sprintf("deck not found [Id: %s]", err.ID)
}

// IsDeckNotFoundError checks if an error is a DeckNotFoundError.
func IsDeckNotFoundError(err error) bool {
	_, ok := err.(DeckNotFoundError)

	return ok
}

// InvalidIDError represents an error that ID is in invalid format.
type InvalidIDError struct {
	ID string
}

// IsInvalidIDError checks if an error is an InvalidIDError.
func IsInvalidIDError(err error) bool {
	_, ok := err.(InvalidIDError)

	return ok
}

// Error returns the error message.
func (err InvalidIDError) Error() string {
	return fmt.Sprintf("invalid id format [id: %s]", err.ID)
}

// DeckRemainingExceededError will be thrown if requests cards of the deck is more than the cards in the deck.
type DeckRemainingExceededError struct {
	Count     int
	Remaining int
}

// Error returns the error message.
func (err DeckRemainingExceededError) Error() string {
	return fmt.Sprintf("the requested number of cards (%d) exceeds the cards in the deck (%d)", err.Count, err.Remaining)
}

// IsDeckRemainingExceededError checks if an error is a DeckRemainingExceededError.
func IsDeckRemainingExceededError(err error) bool {
	_, ok := err.(DeckRemainingExceededError)

	return ok
}
