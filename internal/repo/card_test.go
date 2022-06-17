package repo_test

import (
	"testing"

	"scenario/internal/repo"

	"github.com/stretchr/testify/assert"
)

func TestAllCards(t *testing.T) {
	allCards := repo.AllCard()

	assert.Len(t, allCards, 52)
}

func TestCardsValues(t *testing.T) {
	values := repo.CardsValues()

	assert.Len(t, values, 13)
}

func TestCardsSuits(t *testing.T) {
	suits := repo.CardsSuits()

	assert.Len(t, suits, 4)
}

func TestGetShortValue(t *testing.T) {
	value := repo.GetShortValue("ACE")

	assert.Equal(t, "A", value)
}

func TestGetLongValue(t *testing.T) {
	value := repo.GetLongValue("A")

	assert.Equal(t, "ACE", value)
}

func TestGetLongValueNumber(t *testing.T) {
	value := repo.GetLongValue("10")

	assert.Equal(t, "10", value)
}

func TestGetLongValueInvalid(t *testing.T) {
	value := repo.GetLongValue("I")

	assert.Equal(t, "I", value)
}

func TestGetShortSuit(t *testing.T) {
	suit := repo.GetShortSuit("SPADES")

	assert.Equal(t, "S", suit)
}

func TestGetLongSuit(t *testing.T) {
	suit := repo.GetLongSuit("S")

	assert.Equal(t, "SPADES", suit)
}

func TestGetLongSuitInvalid(t *testing.T) {
	suit := repo.GetLongSuit("I")

	assert.Equal(t, "I", suit)
}
