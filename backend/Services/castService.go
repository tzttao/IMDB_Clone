package Services

import "backend/Database"

type CastFromSearch struct {
	MovieId         int
	CastName        string
	CastId          int
	CastDescription string
	CastImage       string
}

func SearchCastByMovieId(movieId int) []CastFromSearch {
	var result []CastFromSearch
	Database.DB.Raw("SELECT movie_cast.movie_id, cast.* FROM movie_cast INNER JOIN  cast ON movie_cast.cast_id = cast.cast_id WHERE  movie_cast.movie_id = ?", movieId).Scan(&result)
	return result
}
