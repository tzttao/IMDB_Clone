package Models

type Watched_list struct {
	user_id  uint
	movie_id uint
	User     User  `gorm:"ForeignKey:User_id"`  // 设置关联模型并指定UserId作为外键
	Movie    Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键

}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Watched_list) TableName() string {
	return "watched_list"
}
