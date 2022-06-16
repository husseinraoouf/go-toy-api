package models

import "scenario/internal/repo"

// Card represents a card.
// swagger:model card
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func NewCards(cards []*repo.DeckCard) []Card {

	response := make([]Card, len(cards))
	for i, v := range cards {
		response[i] = Card{
			Code:  v.Card.Code,
			Value: v.Card.Value,
			Suit:  v.Card.Suit,
		}
	}

	return response
}
