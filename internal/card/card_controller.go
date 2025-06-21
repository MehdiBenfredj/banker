package card

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type CardController struct {
	CardService *CardService
}

func NewCardController(service *CardService) *CardController {
	return &CardController{CardService: service}
}

func NewCardModule(db *sql.DB) *CardController {
	repo := NewCardRepository(db)
	svc := NewCardService(repo)
	return NewCardController(svc)
}

func (cardController *CardController) Route(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		if request.URL.Query().Has("user_id") {
			user_id := request.URL.Query().Get("user_id")
			cards, error := cardController.CardService.GetCardsByUserID(user_id)
			if error != nil {
				http.Error(writer, "Could not get cards of user : " + user_id, http.StatusInternalServerError)
				return 
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(cards)
		} else if request.URL.Query().Has("card_id") {
			card_id := request.URL.Query().Get("card_id")
			card, error := cardController.CardService.GetCardByID(card_id)
			if error != nil {
				http.Error(writer, "Could not get card with card_id : " + card_id, http.StatusInternalServerError)
				return 
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(card)
		} else {
			cards, error := cardController.CardService.GetAllCards()
			if error != nil {
				http.Error(writer, "Could not get cards", http.StatusInternalServerError)
				return 
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(cards)
		}
	case http.MethodPost:
		var card Card
		error := json.NewDecoder(request.Body).Decode(&card)
		if error != nil {
			http.Error(writer, "Could not decode data", http.StatusInternalServerError)
			return
		}
		if error = cardController.CardService.CreateCard(card.UserID, card.AccountID); error != nil {
			http.Error(writer, "Could not create card", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Card created successfully!"))
	case http.MethodDelete:
		if !request.URL.Query().Has("card_id") {
			http.Error(writer, "Missing card_id", http.StatusBadRequest)
			return
		}
		card_id := request.URL.Query().Get("card_id")
		if error := cardController.CardService.DeleteCard(card_id); error != nil {
			http.Error(writer, "Could not delete card : " + card_id, http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Card deleted successfuly"))
		
	}
}