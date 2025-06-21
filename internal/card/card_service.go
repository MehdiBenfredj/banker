package card

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bvnk/ccgen"
)

type CardService struct {
	CardRepository *CardRepository
}

func (cardService *CardService) DeleteCard(card_id string) error {
	return cardService.CardRepository.DeleteCard(card_id)
}

func (cardService *CardService) CreateCard(user_id string, account_id string) error {	
	cardNumber := ccgen.CardType.Generate(ccgen.Visa)
	cvv := strconv.Itoa(rand.Intn(900) + 100)
	expiration := time.Now().AddDate(2, 0, 0)
	limit := rand.Intn(4000)
	return cardService.CardRepository.CreateCard(user_id, account_id, cardNumber, expiration, cvv, limit)
}

func (cardService *CardService) GetAllCards() ([]Card, error) {
	return cardService.CardRepository.GetAllCards()	
}

func (cardService *CardService) GetCardsByUserID(user_id string) ([]Card, error) {
	return cardService.CardRepository.GetCardsByUserID(user_id)
}

func (cardService *CardService) GetCardByID(card_id string) (*Card, error) {
	return cardService.CardRepository.GetCardByID(card_id)
}

func NewCardService(repository *CardRepository) *CardService {
	return &CardService{CardRepository: repository}
}
