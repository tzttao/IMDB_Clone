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
}

// LogIn Sign In
func LogIn(user User) error {
	err := Database.DB.Where("username = ? AND password = ?", user.Username, user.Password).Find(&user).Error
	fmt.Println(err)
	return err
}

//Signup
func SignUp(user User) int64 {
	user = User{Username: user.Username, Password: user.Password, User_type: 0, Gender: user.Gender, Age: user.Age}
	result := Database.DB.Create(&user) // 通过数据的指针来创建
	affected := result.RowsAffected     // 返回插入记录的条数
	return affected
}

// DeleteUser Delete
func DeleteUser() {
	Database.DB.Delete(&User{}, 111)
}

// UpdateUser Update
func UpdateUser() {
	Database.DB.Model(&User{}).Update("user_name", "hello")
}

// GetList Query user lists
func GetList(page int, pageSize int) []User {
	var user []User
	Database.DB.Where("status = ? AND user_id = ?", "1", "1").Preload("User").Limit(pageSize).Offset((page - 1) * pageSize).Order("create_time desc").Find(&user)
	return user
}
