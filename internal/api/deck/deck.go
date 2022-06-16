package deck

import (
	"net/http"

	"scenario/internal/context"
	"scenario/internal/models"
	"scenario/internal/service"

	"github.com/ggicci/httpin"
)

// CreateDeckInput is input for creating a deck.
type CreateDeckInput struct {
	Cards    string `in:"query=cards"`
	Shuffled bool   `in:"query=shuffled;default=false"`
}

// CreateDeck creates a deck.
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

	ctx := context.NewContext(w, r)
	input := r.Context().Value(httpin.Input).(*CreateDeckInput)

	deck, err := service.CreateDeck(service.CreateDeckOptions{
		Shuffled: input.Shuffled,
		Cards:    input.Cards,
	})
	if err != nil {
		if models.IsInvalidCardCodeError(err) ||
			models.IsDuplicateCardCodeError(err) {
			ctx.Error(http.StatusUnprocessableEntity, err)
		} else {
			ctx.Error(http.StatusInternalServerError, err)
		}

		return
	}

	ctx.JSON(http.StatusCreated, models.NewDeck(deck))
}

// OpenDeckInput is input for opening a deck.
type OpenDeckInput struct {
	ID string `in:"path=id"`
}

// CreateDeck opens a deck.
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

	ctx := context.NewContext(w, r)

	input := r.Context().Value(httpin.Input).(*OpenDeckInput)

	deck, err := service.GetDeckByID(input.ID)
	if err != nil {
		if models.IsInvalidIDError(err) {
			ctx.Error(http.StatusUnprocessableEntity, err)
		} else if models.IsDeckNotFoundError(err) {
			ctx.Error(http.StatusNotFound, err)
		} else {
			ctx.Error(http.StatusInternalServerError, err)
		}

		return
	}

	ctx.JSON(http.StatusOK, models.NewDeckWithCards(deck))
}

// DrawCardInput is input for drawing cards from a deck.
type DrawCardInput struct {
	ID    string `in:"path=id"`
	Count int    `in:"query=count;default=1"`
}

// CreateDeck draws cards from a deck.
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
	//   '404':
	//     description: not found
	//     schema:
	//       "$ref": "#/definitions/APINotFound"
	//
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/APIError"

	ctx := context.NewContext(w, r)
	input := r.Context().Value(httpin.Input).(*DrawCardInput)

	cards, err := service.DrawFromDeckByID(input.ID, input.Count)
	if err != nil {
		if models.IsInvalidIDError(err) ||
			models.IsDeckRemainingExceededError(err) {
			ctx.Error(http.StatusUnprocessableEntity, err)
		} else if models.IsDeckNotFoundError(err) {
			ctx.Error(http.StatusNotFound, err)
		} else {
			ctx.Error(http.StatusInternalServerError, err)
		}

		return
	}

	ctx.JSON(http.StatusOK, models.NewCards(cards))
}
