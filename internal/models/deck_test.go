package models_test

import (
	"reflect"
	"testing"

	"scenario/internal/models"
	"scenario/internal/repo"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := &repo.Deck{
		ID:        uuid.Nil.String(),
		Shuffled:  true,
		Remaining: 5,
	}

	m1 := models.NewDeck(deck)

	m2 := &models.Deck{
		DeckID:    uuid.Nil.String(),
		Shuffled:  true,
		Remaining: 5,
	}

	assert.True(t, reflect.DeepEqual(m1, m2))
}

func TestNewDeckWithCards(t *testing.T) {
	deck := &repo.Deck{
		ID:        uuid.Nil.String(),
		Shuffled:  true,
		Remaining: 5,
		Cards: []*repo.DeckCard{
			{
				CardCode: "AS",
				Order:    1,
				Card: repo.Card{
					Code:  "AS",
					Value: "ACE",
					Suit:  "SPADES",
				},
			},
			{
				CardCode: "2S",
				Order:    2,
				Card: repo.Card{
					Code:  "2S",
					Value: "2",
					Suit:  "SPADES",
				},
			},
		},
	}

	m1 := models.NewDeckWithCards(deck)

	m2 := &models.DeckWithCards{
		DeckID:    uuid.Nil.String(),
		Shuffled:  true,
		Remaining: 5,
		Cards: []models.Card{
			{
				Code:  "AS",
				Value: "ACE",
				Suit:  "SPADES",
			},
			{
				Code:  "2S",
				Value: "2",
				Suit:  "SPADES",
			},
		},
	}

	assert.True(t, reflect.DeepEqual(m1, m2))
}
