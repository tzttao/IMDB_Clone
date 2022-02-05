package Models

type Comment struct {
	Comment_id          uint `gorm:"primary_key"` // 设置id为主键
	User_id             uint
	Movie_id            uint
	One_rate            float64
	Comment_description string
	User                User  `gorm:"ForeignKey:User_id"`  // 设置关联模型并指定UserId作为外键
	Movie               Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键

}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Comment) TableName() string {
	return "comment"
}
