package Models

type User struct {
	Id        uint   `gorm:"primary_key"` // 设置id为主键
	User_name string // 对应的列名为 `user_name`
	Age       string
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (User) TableName() string {
	return "user"
}
