package Models

type Genre_list struct {
	Genre_id uint `gorm:"primary_key"`
	Genre    string
	//Movie    Movie `gorm:"many2many:movie_genre;"`
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Genre_list) TableName() string {
	return "genre_list"
}
