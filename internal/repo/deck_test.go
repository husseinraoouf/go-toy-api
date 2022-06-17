package repo_test

import (
	"testing"

	"scenario/internal/repo"

	"github.com/stretchr/testify/assert"
)

func TestBeforeCreate(t *testing.T) {
	deck := &repo.Deck{
		Shuffled:  true,
		Remaining: 5,
	}

	err := deck.BeforeCreate(repo.GetDatabase())

	assert.Nil(t, err)
	assert.Equal(t, "52fdfc07-2182-454f-963f-5f0f9a621d72", deck.ID)
}
