package Models

import "backend/Database"

type Cast struct {
	Cast_id          uint `gorm:"primary_key"`
	Cast_name        string
	Cast_description string
	Cast_image       string
}

func AddCast(cast Cast) int64 {
	result := Database.DB.Create(&cast) // 通过数据的指针来创建
	affected := result.RowsAffected     // 返回插入记录的条数
	return affected
}

//update review
func UpdateCast(cast Cast) int64 {
	result := Database.DB.Model(&cast).Updates(&cast)
	//result := Database.DB.Save(&user)
	return result.RowsAffected
}

func DeleteCast(Cast_id int) int64 {
	result := Database.DB.Delete(&Cast{}, Cast_id)
	return result.RowsAffected
}

func SearchCastByName(castName string) []Cast {
	var result []Cast
	//Database.DB.First(&cast, castName)
	Database.DB.Where("cast_name LIKE ?", "%"+castName+"%").Find(&result)
	return result
}

//type Result struct {
//	MovieId         int
//	MovieName       string
//	Year            int
//	Grade           float64
//	Description     string
//	Image           string
//	CastName        string
//	CastId          int
//	CastDescription string
//	CastImage       string
//}

func SearchCastById(castId int) Result {
	var result Result
	Database.DB.Raw("SELECT movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image, cast.cast_name,cast.cast_id, cast.cast_description, cast.cast_image FROM cast INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_id = ? ORDER BY cast.cast_id", castId).Scan(&result)
	return result
}

type CastWithName struct {
	Movie_id   int
	Movie_name string
	Image      string
}

func SearchRelativeMovieByCastId(castId int) []CastWithName {
	var result []CastWithName
	Database.DB.Raw("SELECT movie.movie_name,movie.image,movie.movie_id FROM movie_cast INNER JOIN cast ON movie_cast.cast_id = cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_id = ?", castId).Scan(&result)
	return result
}

//
//func SearchMovieByName(name string) []Movie {
//	var movies []Movie
//	Database.DB.Where("movie_name LIKE ?", "%"+name+"%").Find(&movies)
//	return movies
//}
