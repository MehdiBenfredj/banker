package card

import (
	"database/sql"
)

type CardController struct {
	Service *CardService
}

func NewCardController(service *CardService) *CardController {
	return &CardController{Service: service}
}

func NewCardModule(db *sql.DB) *CardController {
	repo := NewCardRepository(db)
	svc := NewCardService(repo)
	return NewCardController(svc)
}
