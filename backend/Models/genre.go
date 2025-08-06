package Models

import "backend/Database"

type Genre struct {
	Genre_id   int `gorm:"primary_key"`
	Genre_name string
}

func AddGenre(genre Genre) int64 {
	result := Database.DB.Create(&genre) // 通过数据的指针来创建
	affected := result.RowsAffected      // 返回插入记录的条数
	return affected
}

//update review
func UpdateGenre(genre Genre) int64 {
	result := Database.DB.Model(&genre).Updates(&genre)
	//result := Database.DB.Save(&user)
	return result.RowsAffected
}

func DeleteGenre(Genre_id int) int64 {
	result := Database.DB.Delete(&Genre{}, Genre_id)
	return result.RowsAffected
}

//When add new movies, administrator can find whether the genere exists
func SearchGenreName(genreName string) []Genre {
	var genre []Genre
	//Database.DB.First(&cast, castName)
	Database.DB.Where("cast_name LIKE ?", "%"+genreName+"%").Find(&genre)
	return genre
}

//When search the specific genre, users can see the relative movies

type ResultGenre struct {
	GenreName   string
	GenreId     int
	MovieId     int
	MovieName   string
	Year        int
	Grade       float64
	Description string
	Image       string
	//CastName        string
	//CastId          int
	//CastDescription string
	//CastImage       string
}

//func SearchGenreByName(Genre_name string) []ResultGenre {
//	var result []ResultGenre
//	Database.DB.Raw("SELECT genre.GenreName, genre.GenreId, movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image FROM genre INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_name = ? ORDER BY cast.cast_id", castName).Scan(&result)
//	return result
//}

func SearchGenre(genreName string) []ResultGenre {
	var result []ResultGenre
	//Database.DB.Raw("SELECT movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image, cast.cast_name,cast.cast_id, cast.cast_description, cast.cast_image FROM cast INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_name = ? ORDER BY cast.cast_id", castName).Scan(&result)
	Database.DB.Raw("SELECT genre.genre_name,genre.genre_id,movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image FROM genre INNER JOIN movie_genre ON genre.genre_id=movie_genre.genre_id INNER JOIN movie ON movie.movie_id=movie_cast.movie_id WHERE genre.genre_name=?", genreName).Scan(&result)
	return result
}

type GenreWithMovie struct {
	Genre string
}

func SearchMovieWithGenre(movieId int) []GenreWithMovie {
	var result []GenreWithMovie
	//Database.DB.Raw("SELECT movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image, cast.cast_name,cast.cast_id, cast.cast_description, cast.cast_image FROM cast INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_name = ? ORDER BY cast.cast_id", castName).Scan(&result)
	Database.DB.Raw("SELECT genre.genre FROM genre INNER JOIN movie_genre ON genre.genre_id = movie_genre.genre_id INNER JOIN movie ON movie_genre.movie_id = movie.movie_id WHERE movie.movie_id = ?", movieId).Scan(&result)
	return result
}
