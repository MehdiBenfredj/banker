package card

import (
	"database/sql"
	"time"
)

type CardRepository struct {
	db *sql.DB
}

func (cardRepository *CardRepository) CreateCard(user_id string, account_id string, cardNumber string, expiration time.Time, cvv string, limit int) error {
	query := "INSERT INTO cards (user_id, account_id, card_number, expiration, cvv, limit) VALUES ($1, $2, $3, $4, $5, $6)"
	_, error := cardRepository.db.Exec(query, user_id, account_id, cardNumber, expiration, cvv, limit)
	return error
}

func (cardRepository *CardRepository) GetCardByID(card_id string) (*Card, error) {
	query := "SELECT * FROM cards WHERE card_id = $1"
	row := cardRepository.db.QueryRow(query, card_id)
	var card Card
	if error := row.Scan(&card.CardID, &card.UserID, &card.AccountID, &card.CardNumber, &card.Expiration, &card.CVV, &card.Limit); error != nil {
		return nil, error
	}
	return &card, nil

}

func (cardRepository *CardRepository) GetCardsByUserID(user_id string) ([]Card, error) {
	query := "SELECT * FROM cards WHERE user_id = $1"
	rows, error := cardRepository.db.Query(query, user_id)
	if error != nil {
		return nil, error
	}
	rows.Close()
	var cards []Card
	for rows.Next() {
		var card Card
		error = rows.Scan(&card.CardID, &card.UserID, &card.AccountID, &card.CardNumber, &card.Expiration, &card.CVV, &card.Limit)
		if error != nil {
			return nil, error
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (cardRepository *CardRepository) GetAllCards() ([]Card, error) {
	query := "SELECT * FROM cards"
	rows, error := cardRepository.db.Query(query)
	if error != nil {
		return nil, error
	}
	rows.Close()
	var cards []Card
	for rows.Next() {
		var card Card
		error = rows.Scan(&card.CardID, &card.UserID, &card.AccountID, &card.CardNumber, &card.Expiration, &card.CVV, &card.Limit)
		if error != nil {
			return nil, error
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (cardRepository *CardRepository) DeleteCard(card_id string) error {
	query := "DELETE * FROM cards WHERE  card_id = $1"
	_, error := cardRepository.db.Exec(query, card_id)
	return error
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{db: db}
}
