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

func (s *UserService) GetUserByID(id int) (User, error) {
	user, err := s.Repository.GetUserByID(id)
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

func (s *UserService) UpdateUser(id int, firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return s.Repository.UpdateUser(id, firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repository.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	users, err := s.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}