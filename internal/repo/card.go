package repo

import (
	"gorm.io/gorm/clause"
)

// Card represents a card in the database.
type Card struct {
	Code  string `gorm:"type:VARCHAR(3);primaryKey"`
	Value string
	Suit  string

	Decks []*DeckCard `gorm:"foreignKey:CardCode"`
}

const countOfAllCards = 52

var allCards []*Card

// AllCard returns a slice with all possible cards.
func AllCard() []*Card {
	return allCards
}

// CardsValues returns a slice with all possible card values.
func CardsValues() []string {
	return []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
}

// CardsSuits returns a slice with all possible card suits.
func CardsSuits() []string {
	return []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
}

func init() {
	initAllCards()
	RegisterModel(new(Card), func() error {
		result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(allCards)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

// GetShortValue gets the short card value from the long value.
// "ACE" -> "A".
func GetShortValue(value string) string {
	valuesMap := map[string]string{
		"ACE":   "A",
		"JACK":  "J",
		"QUEEN": "Q",
		"KING":  "K",
	}

	if v, exist := valuesMap[value]; exist {
		return v
	}

	return value
}

// GetLongValue gets the long card value from the short value.
// "A" -> "ACE".
func GetLongValue(value string) string {
	valuesMap := map[string]string{
		"A": "ACE",
		"J": "JACK",
		"Q": "QUEEN",
		"K": "KING",
	}

	if v, exist := valuesMap[value]; exist {
		return v
	}

	return value
}

// GetShortSuit gets the short card suit from the long suit.
// "SPADES" -> "S".
func GetShortSuit(suit string) string {
	return suit[0:1]
}

// GetLongSuit gets the long card suit from the short suit.
// "S" -> "SPADES".
func GetLongSuit(suit string) string {
	suitsMap := map[string]string{
		"S": "SPADES",
		"D": "DIAMONDS",
		"C": "CLUBS",
		"H": "HEARTS",
	}

	if v, exist := suitsMap[suit]; exist {
		return v
	}

	return suit
}

// initAllCards generates all possible cards and put it in global variable `allCards`.
func initAllCards() {
	values := CardsValues()
	suits := CardsSuits()

	allCards = make([]*Card, countOfAllCards)

	for i, suit := range suits {
		for j, value := range values {
			code := GetShortValue(value) + GetShortSuit(suit)

			allCards[i*13+j] = &Card{
				Code:  code,
				Value: value,
				Suit:  suit,
			}
		}
	}
}
