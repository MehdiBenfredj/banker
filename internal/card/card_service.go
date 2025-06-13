package card 

type CardService struct {
	Repository *CardRepository
}

func NewCardService(repository *CardRepository) *CardService {
	return &CardService{Repository: repository}
}