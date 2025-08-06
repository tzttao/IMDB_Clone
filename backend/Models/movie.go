package Models

import "backend/Database"

type Movie struct {
	Movie_id    int `gorm:"primary_key"`
	Movie_name  string
	Year        int
	Grade       float64
	Description string
	Image       string
	//Genre_list []*Genre_list `gorm:"many2many:movie_genre;"`
}

// AddMovie add movie but do not add the grade
func AddMovie(movie Movie) int64 {
	movie = Movie{
		Movie_id:    movie.Movie_id,
		Movie_name:  movie.Movie_name,
		Year:        movie.Year,
		Description: movie.Description}
	result := Database.DB.Create(&movie)
	return result.RowsAffected
}

//update movie
func UpdateMovie(movie Movie) int64 {
	result := Database.DB.Model(&movie).Updates(movie)
	return result.RowsAffected
}

//delete movie
func DeleteMovie(movieId int) int64 {
	result := Database.DB.Delete(&Movie{}, movieId)
	return result.RowsAffected
}

//search movie by movie id
func SearchMovieByMovieId(movieId int) Movie {
	var movie Movie
	Database.DB.First(&movie, movieId)
	return movie
}

func SearchMovieByName(name string) []Movie {
	var movies []Movie
	Database.DB.Where("movie_name LIKE ?", "%"+name+"%").Find(&movies)
	return movies
}

func SearchMovieByYear(year int) []Movie {
	var movies []Movie
	Database.DB.Where("year = ?", year).Find(&movies)
	return movies
}

func TopMovie() []Movie {
	var movies []Movie
	Database.DB.Raw("select * from movie ORDER BY grade desc LIMIT 3").Scan(&movies)
	return movies
}

type Result struct {
	MovieId         int
	MovieName       string
	Year            int
	Grade           float64
	Description     string
	Image           string
	CastName        string
	CastId          int
	CastDescription string
	CastImage       string
}

func SearchMovieByCast(castName string) []Result {
	var result []Result
	Database.DB.Raw("SELECT movie.movie_id, movie.movie_name, movie.year, movie.grade, movie.description, movie.Image, cast.cast_name,cast.cast_id, cast.cast_description, cast.cast_image FROM cast INNER JOIN movie_cast ON cast.cast_id = movie_cast.cast_id INNER JOIN movie ON movie_cast.movie_id = movie.movie_id WHERE cast.cast_name = ? ORDER BY cast.cast_id", castName).Scan(&result)
	return result
}
