package main

import "math/rand"

type Account struct {
	FirstName string
	LastName  string
	ID        int
	Number    int64
	Balance   int64
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
