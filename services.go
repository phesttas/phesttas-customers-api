package main

type ServiceInterface interface {
	Create()
	Update()
	Delete()
	Find()
	FindOne()
}

type UserService struct {
}

func (us UserService) Create(UserDTORequest) error {
	return nil
}

func (us UserService) Update(UserDTORequest) error {
	return nil

}

func (us UserService) Delete(UserDTORequest) error {
	return nil

}

func (us UserService) Find(UserDTORequest) ([]UserDTOResponse, error) {
	return []UserDTOResponse{}, nil
}

func (us UserService) FindOne(UserDTORequest) (UserDTOResponse, error) {
	return UserDTOResponse{}, nil
}
