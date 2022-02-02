package Services

import (
	Database "backend/Database" // 引入数据库全局变量所在的包
	Models "backend/Models"
	"fmt"
)

//add data
func AddUser() *Models.User {
	var user1 = new(Models.User)
	user := Models.User{Id: 111, User_name: "hahahha", Age: "20"}

	result := Database.Db.Create(&user) // pass pointer of data to Create

	fmt.Println(user.Id)             // returns inserted data's primary key
	fmt.Println(result.Error)        // returns error
	fmt.Println(result.RowsAffected) // returns inserted records count
	return user1
}

//delete
func DeleteUser() {
	Database.Db.Delete(&Models.User{}, 111)
}

//update

func UpdateUser() {
	Database.Db.Model(&Models.User{}).Update("user_name", "hello")
}

// 查询单条数据
func GetOne() *Models.User {
	var user = new(Models.User)
	Database.Db.Where("id = ?", "1").Preload("User").First(&user)
	fmt.Println(user.User_name)
	return user
}

// 查询列表数据
func GetList(page int, pageSize int) []Models.User {
	var user []Models.User
	Database.Db.Where("status = ? AND user_id = ?", "1", "1").Preload("User").Limit(pageSize).Offset((page - 1) * pageSize).Order("create_time desc").Find(&user)
	return user
}
