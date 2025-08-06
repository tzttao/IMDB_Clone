package Models

import "backend/Database"

type Movie_genre struct {
	Movie_id int
	Genre_id int
	Movie    Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键
	Genre    Genre `gorm:"ForeignKey:Genre_id"` // 设置关联模型并指定UserId作为外键
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
//func (Movie_genre) TableName() string {
//	return "movie_genre"
//}

func AddMovieGenre(movieGenre Movie_genre) int64 {
	result := Database.DB.Create(&movieGenre) // 通过数据的指针来创建
	affected := result.RowsAffected           // 返回插入记录的条数
	return affected
}
