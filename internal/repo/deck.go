package repo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Deck represents a deck in the database.
type Deck struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Shuffled  bool
	Remaining int       `gorm:"default:52"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Cards []*DeckCard `gorm:"foreignKey:DeckID"`
}

// DeckCard represents a deck relation to card in the database.
type DeckCard struct {
	DeckID   string `gorm:"primaryKey"`
	CardCode string `gorm:"primaryKey"`
	Order    int

	Deck Deck `gorm:"foreignKey:DeckID"`
	Card Card `gorm:"foreignKey:CardCode"`
}

func init() {
	RegisterModel(new(Deck))
	RegisterModel(new(DeckCard))
}

func (u *Deck) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()

	return nil
}
