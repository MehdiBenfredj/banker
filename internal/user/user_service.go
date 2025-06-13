package user

type UserService struct {
	Repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{Repository: repository}
}


//CRUD operations
func (s *UserService) CreateUser(firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return s.Repository.CreateUser(firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (s *UserService) GetUserByID(user_id string) (User, error) {
	user, err := s.Repository.GetUserByID(user_id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByLastName(lastName string) ([]User, error) {
	users, err := s.Repository.GetUserByLastName(lastName)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(user_id string, firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return s.Repository.UpdateUser(user_id, firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (s *UserService) DeleteUser(user_id string) error {
	return s.Repository.DeleteUser(user_id)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	users, err := s.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}