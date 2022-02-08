package Models

type Movie struct {
	Movie_id   uint `gorm:"primary_key"`
	Movie_name string
	Year       uint
	Rate       float64
	Desciption string
	//Genre_list []*Genre_list `gorm:"many2many:movie_genre;"`
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Movie) TableName() string {
	return "movie"
}
