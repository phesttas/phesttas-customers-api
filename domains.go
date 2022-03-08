package main

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTORequest struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type UserDTOResponse struct {
	Id       string
	Name     string
	Email    string
	Password string
}
