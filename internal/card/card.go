package card

import "time"

type Card struct {
	CardID     string    `json:"card_id"`
	UserID     string    `json:"user_id"`
	AccountID  string    `json:"account_id"`
	CardNumber string    `json:"card_number"`
	Expiration time.Time `json:"expiration"`
	CVV        string    `json:"cvv"`
	Limit      int       `json:"limit"`
}
