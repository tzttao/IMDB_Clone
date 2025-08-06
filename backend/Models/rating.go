package Models

import (
	"backend/Database"
)

type Rating struct {
	User_id  int
	Movie_id int
	Score    float64
}

//func (Rating) TableName() string {
//	return "rating"
//}

func AddRating(rating Rating) int64 {
	result := Database.DB.Create(&rating) // 通过数据的指针来创建
	affected := result.RowsAffected       // 返回插入记录的条数
	return affected
}

//update review
func UpdateRating(rating Rating) int64 {
	//result := Database.DB.Model(&rating).Updates(Rating{User_id: rating.User_id, Movie_id: rating.Movie_id})
	//result := Database.DB.Save(&user)

	result := Database.DB.Exec("update rating set rating.score=? where rating.user_id=? and rating.movie_id=?", rating.Score, rating.User_id, rating.Movie_id)

	return result.RowsAffected
}

//func DeleteRating(Review_id int) int64 {
//	result := Database.DB.Delete(&Review{}, Review_id)
//	return result.RowsAffected
//}

func DeleteRating(rating Rating) int64 {
	rating = Rating{User_id: rating.User_id, Movie_id: rating.Movie_id, Score: rating.Score}
	result := Database.DB.Delete(&rating)
	return result.RowsAffected
}

func ReadRating(UserId, MovieId int) Rating {
	var rating Rating
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Where("user_id = ? and movie_id = ?", UserId, MovieId).Find(&rating)
	Database.DB.Exec("update movie set movie.grade= (select CAST(AVG(rating.score) as decimal(10,1)) from rating where rating.movie_id=?) where movie.movie_id=?", MovieId, MovieId)
	return rating
}

func UpdateGrade(movie_id int) int64 {
	//var result []Result
	//Database.DB.Raw("SELECT movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image, cast.cast_name,cast.cast_id, cast.cast_description, cast.cast_image FROM cast INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_name = ? ORDER BY cast.cast_id", castName).Scan(&result)

	result := Database.DB.Exec("update movie set movie.grade= (select AVG(rating.score) from rating where rating.movie_id=?) where movie.movie_id=?", movie_id, movie_id)
	//db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
	//	sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
	return result.RowsAffected
}
