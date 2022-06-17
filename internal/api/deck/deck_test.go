package deck_test

import (
	"net/http"
	"scenario/internal/testutils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {

	server := testutils.NewTestServer()

	w := server.Post("/deck")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"deck_id":"52fdfc07-2182-454f-963f-5f0f9a621d72","shuffled":false,"remaining":52}`, strings.TrimSpace(w.Body.String()))
}

func TestCreateDeckInvalidCardCode(t *testing.T) {

	server := testutils.NewTestServer()

	w := server.Post("/deck?cards=AS,5H,10Y,7S")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, `{"message":"invalid card code [code: 10Y]"}`, strings.TrimSpace(w.Body.String()))
}

func TestCreateDeckDuplicateCardCode(t *testing.T) {

	server := testutils.NewTestServer()

	w := server.Post("/deck?cards=AS,5H,AS,7S")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, `{"message":"duplicate card code [code: AS]"}`, strings.TrimSpace(w.Body.String()))
}

func TestCreateDeckInvalidShuffledValue(t *testing.T) {

	server := testutils.NewTestServer()

	w := server.Post("/deck?shuffled=invalid")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, `{"field":"Shuffled","source":"query","value":"invalid","error":"strconv.ParseBool: parsing \"invalid\": invalid syntax"}`, strings.TrimSpace(w.Body.String()))
}
