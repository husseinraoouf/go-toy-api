package service

import (
	"errors"
	"fmt"
	"math/rand"
	"scenario/internal/models"
	"scenario/internal/repo"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

// CreateDeckOptions holds options to create a Deck
type CreateDeckOptions struct {
	Cards    string
	Shuffled bool
}

// CreateDeck creates a Deck
func CreateDeck(opts CreateDeckOptions) (*repo.Deck, error) {
	allCards := repo.AllCard()

	var cards []*repo.Card

	if opts.Cards == "" {
		// if cards filter is empty we put all the cards in the deck
		cards = allCards
	} else {
		cardsFilter := strings.Split(opts.Cards, ",")

		if err := ValidateCards(cardsFilter); err != nil {
			return nil, err
		}

		cards = make([]*repo.Card, len(cardsFilter))
		// I need to make the counter only increase when
		i := 0
		for _, c := range allCards {
			if slices.Contains(cardsFilter, c.Code) {
				cards[i] = c
				i++
			}
		}
	}

	if opts.Shuffled {
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	db := repo.GetDatabase()

	deck := &repo.Deck{Shuffled: opts.Shuffled, Remaining: len(cards), Cards: make([]*repo.DeckCard, len(cards))}
	for i, c := range cards {
		deck.Cards[i] = &repo.DeckCard{
			DeckId:   deck.Id,
			CardCode: c.Code,
			Order:    i + 1,
		}
	}

	result := db.Create(deck)
	if result.Error != nil {
		return nil, fmt.Errorf("insert deck: %v", result.Error)
	}

	return deck, nil
}

// GetDeck returns the deck with the given id.
func GetDeckById(id string) (*repo.Deck, error) {

	if err := ValidateId(id); err != nil {
		return nil, err
	}

	db := repo.GetDatabase()

	deck := new(repo.Deck)

	result := db.Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Order("deck_cards.order ASC")
	},
	).Preload("Cards.Card").Where("id = ?", id).First(deck)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrDeckNotFound{Id: id}
		} else {
			return nil, fmt.Errorf("fetching deck: %v", result.Error)
		}
	}

	return deck, nil
}

// DrawFromDeckById draws cards from the deck. returns the drawn cards.
func DrawFromDeckById(id string, count int) ([]*repo.DeckCard, error) {

	if err := ValidateId(id); err != nil {
		return nil, err
	}

	db := repo.GetDatabase()

	deck := new(repo.Deck)

	result := db.Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Order("deck_cards.order ASC").Limit(count)
	},
	).Preload("Cards.Card").Where("id = ?", id).First(deck)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrDeckNotFound{Id: id}
		} else {
			return nil, fmt.Errorf("fetching deck: %v", result.Error)
		}
	}

	// if the requested number of cards is larger than the cards in the deck
	if deck.Remaining < count || len(deck.Cards) < count {
		return nil, models.ErrDeckRemainingExceeded{Count: count, Remaining: deck.Remaining}
	}

	// update deck and its relations in one transaction to make
	// sure that deck.Remaining is always right
	err := db.Transaction(func(tx *gorm.DB) error {
		// remove cards relations to the deck
		result = tx.Delete(deck.Cards)
		if result.Error != nil {
			return fmt.Errorf("updating deck: %v", result.Error)
		}

		// update deck.Remaining in the deck
		deck.Remaining -= count
		result = tx.Model(deck).Select("Remaining").Updates(deck)
		if result.Error != nil {
			return fmt.Errorf("updating deck: %v", result.Error)
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return nil, err
	}

	return deck.Cards, nil
}

func ValidateCards(cards []string) error {
	frequency_map := make(map[string]bool)

	for _, c := range cards {
		if _, exists := frequency_map[c]; exists {
			return models.ErrDuplicateCardCode{
				CardCode: c,
			}
		}
		frequency_map[c] = true

		if len := len(c); len > 3 || len == 0 {
			return models.ErrInvalidCardCode{
				CardCode: c,
			}
		}

		values := repo.CardsValues()
		suits := repo.CardsSuits()

		value := repo.GetLongValue(c[:len(c)-1])
		suit := repo.GetLongSuit(c[len(c)-1:])

		if !slices.Contains(values, value) || !slices.Contains(suits, suit) {
			return models.ErrInvalidCardCode{
				CardCode: c,
			}
		}

	}
	return nil
}

func ValidateId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return models.ErrInvalidId{
			Id: id,
		}
	}

	return nil
}
