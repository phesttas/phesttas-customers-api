package main

type ControllerInterface interface {
	Create()
	Update()
	Delete()
	Find()
	FindOne()
}

type UserController struct{}

func (uc UserController) Create() {

}

func (uc UserController) Update() {

}

func (uc UserController) Delete() {

}

func (uc UserController) Find() {

}

func (uc UserController) FindOne() {

}
