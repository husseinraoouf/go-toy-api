package repo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deck struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	Shuffled  bool
	Remaining int       `gorm:"default:52"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Cards []*DeckCard `gorm:"foreignKey:DeckId"`
}

type DeckCard struct {
	DeckId   string `gorm:"primaryKey"`
	CardCode string `gorm:"primaryKey"`
	Order    int

	Deck Deck `gorm:"foreignKey:DeckId"`
	Card Card `gorm:"foreignKey:CardCode"`
}

func init() {
	RegisterModel(new(Deck))
	RegisterModel(new(DeckCard))
}

func (u *Deck) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.NewString()
	return nil
}
