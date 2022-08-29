package models

import "time"

type Users struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
	TimeIn    time.Time
	TimeOut   time.Time
}

// type UserType int

// func (UserType) Administrator() UserType {
// 	return 1
// }
// func (UserType) Standard() UserType {
// 	return 2
// }
