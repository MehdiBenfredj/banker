package account

type Account struct {
	AccountID string `json:"account_id"`
	UserID    string `json:"user_id"`
	BIC       string `json:"bic"`
	IBAN      string `json:"iban"`
}
