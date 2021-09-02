package users

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u User) DisplayName() string {
	return u.FirstName + " " + u.LastName
}
