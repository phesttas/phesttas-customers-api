package main

type ControllerInterface interface {
	Create()
	Update()
	Delete()
	Find()
	FindOne()
}

type UserController struct{}

func (uc UserController) Create(UserDTORequest) error {
	return nil
}

func (uc UserController) Update(UserDTORequest) error {
	return nil

}

func (uc UserController) Delete(UserDTORequest) error {
	return nil
}

func (uc UserController) Find(UserDTORequest) ([]UserDTOResponse, error) {
	return nil, nil
}

func (uc UserController) FindOne(UserDTORequest) (UserDTOResponse, error) {
	return UserDTOResponse{
		Id:       "1",
		Name:     "John",
		Email:    "john@mail.com",
		Password: "123456",
	}, nil
}
