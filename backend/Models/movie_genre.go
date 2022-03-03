package Models

type Movie_genre struct {
	Movie_id   uint
	Genre_id   uint
	Movie      Movie      `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键
	Genre_list Genre_list `gorm:"ForeignKey:Genre_id"` // 设置关联模型并指定UserId作为外键
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Movie_genre) TableName() string {
	return "movie_genre"
}
