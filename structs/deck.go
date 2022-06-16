package structs

import "scenario/repo"

// Deck represents a deck.
// swagger:model deck
type Deck struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// DeckWithCards represents an Deck with cards.
// swagger:model deckWithCards
type DeckWithCards struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

func NewDeck(deck *repo.Deck) *Deck {
	return &Deck{
		DeckId:    deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
}

func NewDeckWithCards(deck *repo.Deck) *DeckWithCards {
	response := &DeckWithCards{
		DeckId:    deck.Id,
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
