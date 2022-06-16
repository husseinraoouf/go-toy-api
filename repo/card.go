package repo

var allCards []*Card

func GetAllCard() []*Card {
	return allCards
}

func GetCardsValues() []string {
	return []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
}

func GetCardsSuits() []string {
	return []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
}

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

func GetShortSuit(suit string) string {
	return suit[0:1]
}

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

type Card struct {
	Code  string `gorm:"type:VARCHAR(3);primaryKey"`
	Value string
	Suit  string

	Decks []*DeckCard `gorm:"foreignKey:CardCode"`
}

func init() {
	RegisterModel(new(Card))

	values := GetCardsValues()
	suits := GetCardsSuits()

	allCards = make([]*Card, 52)
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
