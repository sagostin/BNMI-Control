package main

type User struct {
	ID                    int
	AccountID             int
	FirstName             string
	LastName              string
	Email                 string
	RegistrationTimestamp string
	LastLoginTimestamp    string
	PasswordHash          string // sha512 hash?
	Role                  string
}
