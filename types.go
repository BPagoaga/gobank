package main

import "math/rand"

type Account struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        int    `json:"id"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		Balance:   0,
	}
}
