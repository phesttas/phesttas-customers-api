package main

import (
	"fmt"
	"net/http"
)

type EndpointInterface interface {
	Create()
	Update()
	Delete()
	Find()
	FindOne()
}

type UserEndpoint struct{}

func (e UserEndpoint) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createUser")
}

func (e UserEndpoint) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateUser")
}

func (e UserEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete User Endpoint")
}

func (e UserEndpoint) Find(w http.ResponseWriter, r *http.Request) {
	us := UserService{}
	u, err := us.Find(UserDTORequest{})
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, u)
}

func (e UserEndpoint) FindOne(w http.ResponseWriter, r *http.Request) {
	uc := UserService{}
	u, err := uc.FindOne(UserDTORequest{})
	if err != nil {
		fmt.Fprint(w, "error")
	}
	fmt.Fprint(w, u.Name)
}
