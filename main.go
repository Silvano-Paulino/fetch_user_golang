package main

type User struct {
	Id   string
	Name string
}

func FetchUser(db Database, id int) (*User, error) {
	return db.GetUserId(id)
}

type Database interface {
	GetUserId(id int) (*User, error)
}