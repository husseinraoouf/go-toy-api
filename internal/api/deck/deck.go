// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: ElHussein Abdelraouf<hussein@raoufs.me>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package deck

import (
	"net/http"
	"scenario/internal/context"
	"scenario/internal/models"
	"scenario/internal/service"

	"github.com/ggicci/httpin"
)

// CreateDeckInput input for creating a deck
type CreateDeckInput struct {
	Cards    string `in:"query=cards"`
	Shuffled bool   `in:"query=shuffled;default=false"`
}

// CreateDeck create a deck
func CreateDeck(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /deck deck createDeck
	//
	// creates a new Deck
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: cards
	//   in: query
	//   description: cards to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: shuffled
	//   in: query
	//   description: whether the cards in the deck should be shuffled or not
	//   required: false
	//   type: boolean
	// responses:
	//   '201':
	//     description: deck created
	//     schema:
	//       "$ref": "#/definitions/deck"
	//
	//   '422':
	//     description: validation error
	//     schema:
	//       "$ref": "#/definitions/APIValidationError"
	//
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/APIError"

	c := context.NewContext(w, r)
	input := r.Context().Value(httpin.Input).(*CreateDeckInput)

	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: input.Shuffled,
		Cards:    input.Cards,
	})
	if err != nil {
		if models.IsErrInvalidCardCode(err) ||
			models.IsErrDuplicateCardCode(err) {
			c.Error(http.StatusUnprocessableEntity, err)
		} else {
			c.Error(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusCreated, models.NewDeck(deck))
}

// OpenDeckInput input for opening a deck
type OpenDeckInput struct {
	Id string `in:"path=id"`
}

// CreateDeck open a deck
func OpenDeck(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /deck/{id} deck openDeck
	//
	// opens a Deck
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the deck
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: deck with cards response
	//     schema:
	//       "$ref": "#/definitions/deckWithCards"
	//
	//   '422':
	//     description: validation error
	//     schema:
	//       "$ref": "#/definitions/APIValidationError"
	//
	//   '404':
	//     description: not found
	//     schema:
	//       "$ref": "#/definitions/APINotFound"
	//
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/APIError"

	c := context.NewContext(w, r)

	input := r.Context().Value(httpin.Input).(*OpenDeckInput)

	deck, err := service.GetDeckById(input.Id)
	if err != nil {
		if models.IsErrInvalidId(err) {
			c.Error(http.StatusUnprocessableEntity, err)
		} else if models.IsErrDeckNotFound(err) {
			c.Error(http.StatusNotFound, err)
		} else {
			c.Error(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, models.NewDeckWithCards(deck))
}

// DrawCardInput input for drawing cards from a deck
type DrawCardInput struct {
	Id    string `in:"path=id"`
	Count int    `in:"query=count;default=1"`
}

// CreateDeck draw cards from a deck
func DrawCard(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /deck/{id}/draw deck drawDeck
	//
	// draw cards from a deck
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the deck
	//   required: true
	//   type: string
	// - name: count
	//   in: query
	//   description: how many cards to draw from the deck
	//   required: false
	//   type: integer
	// responses:
	//   '200':
	//     description: cards response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/card"
	//
	//   '422':
	//     description: validation error
	//     schema:
	//       "$ref": "#/definitions/APIValidationError"
	//
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/APIError"

	c := context.NewContext(w, r)
	input := r.Context().Value(httpin.Input).(*DrawCardInput)

	cards, err := service.DrawFromDeckById(input.Id, input.Count)
	if err != nil {
		if models.IsErrInvalidId(err) ||
			models.IsErrDeckRemainingExceeded(err) {
			c.Error(http.StatusUnprocessableEntity, err)
		} else if models.IsErrDeckNotFound(err) {
			c.Error(http.StatusNotFound, err)
		} else {
			c.Error(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, models.NewCards(cards))
}
