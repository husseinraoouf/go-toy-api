package deck_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"scenario/internal/models"
	"scenario/internal/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post("/deck")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t,
		`{"deck_id":"52fdfc07-2182-454f-963f-5f0f9a621d72","shuffled":false,"remaining":52}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestCreateDeckWithParams(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post("/deck?cards=AS,5H&shuffled=true")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t,
		`{"deck_id":"9566c74d-1003-4c4d-bbbb-0407d1e2c649","shuffled":true,"remaining":2}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestCreateDeckInvalidCardCode(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post("/deck?cards=AS,5H,10Y,7S")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t,
		`{"message":"invalid card code [code: 10Y]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestCreateDeckDuplicateCardCode(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post("/deck?cards=AS,5H,AS,7S")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t,
		`{"message":"duplicate card code [code: AS]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestCreateDeckInvalidShuffledValue(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post("/deck?shuffled=invalid")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t,
		//nolint:lll
		`{"field":"Shuffled","source":"query","value":"invalid","error":"strconv.ParseBool: parsing \"invalid\": invalid syntax"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestOpenDeck(t *testing.T) {
	server := testutils.NewTestServer()

	// create deck
	createDeckResponse := server.Post("/deck?cards=AS,5H")

	deck := new(models.Deck)
	if err := json.Unmarshal(createDeckResponse.Body.Bytes(), deck); err != nil {
		t.Error(err)
	}

	w := server.Get(fmt.Sprintf("/deck/%s", deck.DeckID))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,
		//nolint:lll
		fmt.Sprintf(`{"deck_id":"%s","shuffled":false,"remaining":2,"cards":[{"value":"ACE","suit":"SPADES","code":"AS"},{"value":"5","suit":"HEARTS","code":"5H"}]}`, deck.DeckID),
		strings.TrimSpace(w.Body.String()),
	)
}

func TestOpenDeckInvalidId(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Get(fmt.Sprintf("/deck/%s", uuid.Invalid.String()))

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t,
		`{"message":"invalid id format [id: Invalid]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestOpenDeckNotFound(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Get(fmt.Sprintf("/deck/%s", uuid.Nil.String()))

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t,
		`{"message":"deck not found [Id: 00000000-0000-0000-0000-000000000000]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestDrawCard(t *testing.T) {
	server := testutils.NewTestServer()

	// create deck
	createDeckResponse := server.Post("/deck?cards=AS,5H")

	deck := new(models.Deck)
	if err := json.Unmarshal(createDeckResponse.Body.Bytes(), deck); err != nil {
		t.Error(err)
	}

	w := server.Post(fmt.Sprintf("/deck/%s/draw", deck.DeckID))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,
		`[{"value":"ACE","suit":"SPADES","code":"AS"}]`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestDrawCardInvalidId(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post(fmt.Sprintf("/deck/%s/draw", uuid.Invalid.String()))

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t,
		`{"message":"invalid id format [id: Invalid]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}

func TestDrawCardNotFound(t *testing.T) {
	server := testutils.NewTestServer()

	w := server.Post(fmt.Sprintf("/deck/%s/draw", uuid.Nil.String()))

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t,
		`{"message":"deck not found [Id: 00000000-0000-0000-0000-000000000000]"}`,
		strings.TrimSpace(w.Body.String()),
	)
}
