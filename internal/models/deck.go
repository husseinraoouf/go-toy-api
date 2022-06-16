package models

import "scenario/internal/repo"

// Deck represents a deck.
// swagger:model deck
type Deck struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// DeckWithCards represents an Deck with cards.
// swagger:model deckWithCards
type DeckWithCards struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

// NewDeck creates new deck from repo struct.
func NewDeck(deck *repo.Deck) *Deck {
	return &Deck{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
}

// NewDeckWithCards creates new deck with cards from repo struct.
func NewDeckWithCards(deck *repo.Deck) *DeckWithCards {
	response := &DeckWithCards{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
		Cards:     make([]Card, len(deck.Cards)),
	}

	for i, v := range deck.Cards {
		response.Cards[i] = Card{
			Code:  v.Card.Code,
			Value: v.Card.Value,
			Suit:  v.Card.Suit,
		}
	}

	return response
}
