package service_test

import (
	"reflect"
	"testing"

	"scenario/internal/models"
	"scenario/internal/repo"
	"scenario/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: false,
		Cards:    "",
	})

	standardDeck := &repo.Deck{
		ID:        deck.ID,
		Shuffled:  false,
		Remaining: 52,
		Cards:     make([]*repo.DeckCard, 52),
	}

	for i, c := range repo.AllCard() {
		standardDeck.Cards[i] = &repo.DeckCard{
			DeckID:   deck.ID,
			CardCode: c.Code,
			Order:    i + 1,
		}
	}

	// set unwanted fields to specific value to exclude from equality
	deck.CreatedAt = standardDeck.CreatedAt
	deck.UpdatedAt = standardDeck.UpdatedAt

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(standardDeck, deck))
}

func TestCreateDeckFilter(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: false,
		Cards:    "AS,5S",
	})

	standardDeck := &repo.Deck{
		ID:        deck.ID,
		Shuffled:  false,
		Remaining: 2,
		Cards: []*repo.DeckCard{
			{
				DeckID:   deck.ID,
				CardCode: "AS",
				Order:    1,
			},
			{
				DeckID:   deck.ID,
				CardCode: "5S",
				Order:    2,
			},
		},
	}

	// set unwanted fields to specific value to exclude from equality
	deck.CreatedAt = standardDeck.CreatedAt
	deck.UpdatedAt = standardDeck.UpdatedAt

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(standardDeck, deck))
}

func TestCreateDeckShuffled(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: true,
		Cards:    "AS,5S,8H",
	})

	standardDeck := &repo.Deck{
		ID:        deck.ID,
		Shuffled:  true,
		Remaining: 3,
		Cards: []*repo.DeckCard{
			{
				DeckID:   deck.ID,
				CardCode: "AS",
				Order:    1,
			},
			{
				DeckID:   deck.ID,
				CardCode: "8H",
				Order:    2,
			},
			{
				DeckID:   deck.ID,
				CardCode: "5S",
				Order:    3,
			},
		},
	}

	// set unwanted fields to specific value to exclude from equality
	deck.CreatedAt = standardDeck.CreatedAt
	deck.UpdatedAt = standardDeck.UpdatedAt

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(standardDeck, deck))
}

func TestCreateDeckInvalidCards(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: true,
		Cards:    "AS,5S,8T",
	})

	assert.Nil(t, deck)
	assert.ErrorIs(t, err, models.InvalidCardCodeError{CardCode: "8T"})
}

func TestGetDeckByID(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: false,
		Cards:    "AS,5S",
	})

	deckWithCards, openErr := service.GetDeckByID(deck.ID)

	standardDeckWithCards := &repo.Deck{
		ID:        deck.ID,
		Shuffled:  false,
		Remaining: 2,
		Cards: []*repo.DeckCard{
			{
				DeckID:   deck.ID,
				CardCode: "AS",
				Order:    1,
				Card: repo.Card{
					Code:  "AS",
					Value: "ACE",
					Suit:  "SPADES",
				},
			},
			{
				DeckID:   deck.ID,
				CardCode: "5S",
				Order:    2,
				Card: repo.Card{
					Code:  "5S",
					Value: "5",
					Suit:  "SPADES",
				},
			},
		},
	}

	// set unwanted fields to specific value to exclude from equality
	deckWithCards.CreatedAt = standardDeckWithCards.CreatedAt
	deckWithCards.UpdatedAt = standardDeckWithCards.UpdatedAt

	assert.Nil(t, err)
	assert.Nil(t, openErr)
	assert.True(t, reflect.DeepEqual(standardDeckWithCards, deckWithCards))
}

func TestGetDeckByIDInvalidID(t *testing.T) {
	deck, err := service.GetDeckByID(uuid.Invalid.String())

	assert.Nil(t, deck)
	assert.ErrorIs(t, err, models.InvalidIDError{ID: uuid.Invalid.String()})
}

func TestGetDeckByIDNotFound(t *testing.T) {
	deck, err := service.GetDeckByID(uuid.Nil.String())

	assert.Nil(t, deck)
	assert.ErrorIs(t, err, models.DeckNotFoundError{ID: uuid.Nil.String()})
}

func TestDrawFromDeckByID(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: false,
		Cards:    "AS,5S",
	})

	cards, openErr := service.DrawFromDeckByID(deck.ID, 1)

	standardCards := []*repo.DeckCard{
		{
			DeckID:   deck.ID,
			CardCode: "AS",
			Order:    1,
			Card: repo.Card{
				Code:  "AS",
				Value: "ACE",
				Suit:  "SPADES",
			},
		},
	}

	assert.Nil(t, err)
	assert.Nil(t, openErr)
	assert.True(t, reflect.DeepEqual(standardCards, cards))
}

func TestDrawFromDeckByIDInvalidID(t *testing.T) {
	deck, err := service.DrawFromDeckByID(uuid.Invalid.String(), 1)

	assert.Nil(t, deck)
	assert.ErrorIs(t, err, models.InvalidIDError{ID: uuid.Invalid.String()})
}

func TestDrawFromDeckByIDNotFound(t *testing.T) {
	deck, err := service.DrawFromDeckByID(uuid.Nil.String(), 1)

	assert.Nil(t, deck)
	assert.ErrorIs(t, err, models.DeckNotFoundError{ID: uuid.Nil.String()})
}

func TestDrawFromDeckByIDExceedRemaining(t *testing.T) {
	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: false,
		Cards:    "AS,5S",
	})

	cards, openErr := service.DrawFromDeckByID(deck.ID, 5)

	assert.Nil(t, err)
	assert.Nil(t, cards)
	assert.ErrorIs(t, openErr, models.DeckRemainingExceededError{Count: 5, Remaining: 2})
}

func TestValidateCards(t *testing.T) {
	err := service.ValidateCards([]string{"AS", "5S"})

	assert.Nil(t, err)
}

func TestValidateCardsInvalidCards(t *testing.T) {
	err := service.ValidateCards([]string{"Invalid", "5T"})

	assert.ErrorIs(t, err, models.InvalidCardCodeError{CardCode: "Invalid"})
}

func TestValidateCardsInvalidCards2(t *testing.T) {
	err := service.ValidateCards([]string{"AS", "5T"})

	assert.ErrorIs(t, err, models.InvalidCardCodeError{CardCode: "5T"})
}

func TestValidateCardsDuplicateCards(t *testing.T) {
	err := service.ValidateCards([]string{"AS", "AS"})

	assert.ErrorIs(t, err, models.DuplicateCardCodeError{CardCode: "AS"})
}

func TestValidateID(t *testing.T) {
	err := service.ValidateID(uuid.Nil.String())

	assert.Nil(t, err)
}

func TestValidateIDInvalid(t *testing.T) {
	err := service.ValidateID(uuid.Invalid.String())

	assert.ErrorIs(t, err, models.InvalidIDError{ID: uuid.Invalid.String()})
}
