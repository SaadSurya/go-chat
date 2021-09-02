package users

import "github.com/saadsurya/go-chat/database"

func FindByUsernamePassword(username string, password string) *User {
	db := database.DBConn
	var user User
	db.Where("username = ? and password = ?", username, password).First(&user)
	return &user
}

func CreateUser(user *User) {
	db := database.DBConn
	db.Create(&user)
}

func GetAllUsers() []User {
	db := database.DBConn
	var users []User
	db.Select([]string{"id", "username", "first_name", "last_name"}).Find(&users)
	return users
}
