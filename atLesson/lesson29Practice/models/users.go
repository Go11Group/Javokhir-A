package models

type User struct {
	UserID    string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}

// {"id":"1cfd07d0-1942-4dd2-84e4-c32e8cd146e2","first_name":"Erica","last_name":"Luney","email":"eluneylp@friendfeed.com","gender":"Female","age":132},
