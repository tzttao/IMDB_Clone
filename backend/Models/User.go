package Models

import (
	"backend/Database"
	"fmt"
)

type User struct {
	User_id   int `gorm:"primary_key"` // 设置id为主键
	Username  string
	Password  string
	User_type int
	Gender    string
	Age       int
	Email     string
}

// LogIn Sign In
func LogIn(user User) User {
	Database.DB.Where("username = ? AND password = ?", user.Username, user.Password).Find(&user)
	fmt.Println("user:%#v\n", user)
	return user
}

//Signup
func SignUp(user User) int64 {
	user = User{Username: user.Username, Password: user.Password, User_type: 0, Gender: user.Gender, Age: user.Age, Email: user.Email}
	result := Database.DB.Create(&user) // 通过数据的指针来创建
	affected := result.RowsAffected     // 返回插入记录的条数
	return affected
}

// UpdateProfile UpdateUser Update
func UpdateProfile(user User) int64 {
	result := Database.DB.Model(&user).Updates(User{Username: user.Username, Password: user.Password, Gender: user.Gender, Age: user.Age, Email: user.Email})
	//result := Database.DB.Save(&user)
	return result.RowsAffected
}

// DeleteUser Delete
func DeleteUser(userId int) int64 {
	result := Database.DB.Delete(&User{}, userId)
	return result.RowsAffected
}

// GetList Query user lists
func GetList(name string) []User {
	var users []User
	Database.DB.Where("username = ?", name).Find(&users)
	fmt.Printf("user:%#v\n", users)
	return users
}

func QueryUserInfoByUserId(userId int) User {
	var user User
	Database.DB.Where("User_id = ?", userId).Find(&user)
	fmt.Println("user:%#v\n", user)
	return user
}
