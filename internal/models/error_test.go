package models_test

import (
	"testing"

	"scenario/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInvalidCardCodeError(t *testing.T) {
	err := models.InvalidCardCodeError{
		CardCode: "AT",
	}

	assert.True(t, models.IsInvalidCardCodeError(err))
	assert.Equal(t, "invalid card code [code: AT]", err.Error())
}

func TestDuplicateCardCodeError(t *testing.T) {
	err := models.DuplicateCardCodeError{
		CardCode: "AS",
	}

	assert.True(t, models.IsDuplicateCardCodeError(err))
	assert.Equal(t, "duplicate card code [code: AS]", err.Error())
}

func TestDeckNotFoundError(t *testing.T) {
	err := models.DeckNotFoundError{
		ID: uuid.Nil.String(),
	}

	assert.True(t, models.IsDeckNotFoundError(err))
	assert.Equal(t, "deck not found [Id: 00000000-0000-0000-0000-000000000000]", err.Error())
}

func TestInvalidIDError(t *testing.T) {
	err := models.InvalidIDError{
		ID: uuid.Invalid.String(),
	}

	assert.True(t, models.IsInvalidIDError(err))
	assert.Equal(t, "invalid id format [id: Invalid]", err.Error())
}

func TestDeckRemainingExceededError(t *testing.T) {
	err := models.DeckRemainingExceededError{
		Count:     10,
		Remaining: 5,
	}

	assert.True(t, models.IsDeckRemainingExceededError(err))
	assert.Equal(t, "the requested number of cards (10) exceeds the cards in the deck (5)", err.Error())
}
