package repo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deck struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Shuffled  bool
	Remaining int       `gorm:"default:52"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Cards []*DeckCard `gorm:"foreignKey:DeckId"`
}

type DeckCard struct {
	DeckID   string `gorm:"primaryKey"`
	CardCode string `gorm:"primaryKey"`
	Order    int

	Deck Deck `gorm:"foreignKey:DeckId"`
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
