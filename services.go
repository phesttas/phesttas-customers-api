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

func (us UserService) Create() {

}

func (us UserService) Update() {

}

func (us UserService) Delete() {

}

func (us UserService) Find() {

}

func (us UserService) FindOne() {

}
