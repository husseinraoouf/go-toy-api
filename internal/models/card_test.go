package models_test

import (
	"reflect"
	"testing"

	"scenario/internal/models"
	"scenario/internal/repo"

	"github.com/stretchr/testify/assert"
)

func TestNewCards(t *testing.T) {
	cards := []*repo.DeckCard{
		{
			Order: 1,
			Card: repo.Card{
				Code:  "AS",
				Value: "ACE",
				Suit:  "SPADES",
			},
		},
		{
			Order: 2,
			Card: repo.Card{
				Code:  "2S",
				Value: "2",
				Suit:  "SPADES",
			},
		},
	}

	m1 := models.NewCards(cards)

	m2 := []models.Card{
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
	}

	assert.True(t, reflect.DeepEqual(m1, m2))
}
