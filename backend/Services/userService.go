package Services

import (
	"backend/Database"
	"backend/Models"
)

// Check unique username
func CheckUsername(username string) bool {
	user := Models.User{}
	isFound := Database.DB.Where("username = ?", username).Find(&user).RecordNotFound()
	return isFound
}
