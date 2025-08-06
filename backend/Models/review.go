package Models

import "backend/Database"

type Review struct {
	Review_id      int `gorm:"primary_key"` // 设置id为主键
	User_id        int
	Movie_id       int
	Review_content string
	//User           User  `gorm:"ForeignKey:User_id"`  // 设置关联模型并指定UserId作为外键
	//Movie          Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
//func (Review) TableName() string {
//	return "review"
//}

// create review
func AddReview(review Review) int64 {
	result := Database.DB.Create(&review) // 通过数据的指针来创建
	affected := result.RowsAffected       // 返回插入记录的条数
	return affected
}

//update review
func UpdateReview(review Review) int64 {
	result := Database.DB.Model(&review).Updates(Review{Review_content: review.Review_content})
	//result := Database.DB.Save(&user)
	return result.RowsAffected
}

func DeleteReview(Review_id int) int64 {
	result := Database.DB.Delete(&Review{}, Review_id)
	return result.RowsAffected
}

func ReadReview(User_id int) []Review {
	var review []Review
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Where("user_id = ?", User_id).Find(&review)
	return review
}

type reviewName struct {
	Review_id      int
	User_id        int
	Movie_id       int
	Review_content string
	Username       string
}

func CountNumOfReview(Movie_id int) int {
	var review []reviewName
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Raw("SELECT review.*, user.username FROM review INNER JOIN user ON review.user_id = user.user_id WHERE review.movie_id = ?", Movie_id).Scan(&review)
	return len(review)
}
func ReadReviewByMovieId(Movie_id, pageNo, pageSize int) []reviewName {
	var review []reviewName
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Raw("SELECT review.*, user.username FROM review INNER JOIN user ON review.user_id = user.user_id WHERE review.movie_id = ? Limit ?,?", Movie_id, (pageNo-1)*pageSize, pageSize).Scan(&review)
	return review
}

//func CountNumOfReview1(Movie_id int) int {
//	var count int
//	var review reviewName
//	Database.DB.Raw("SELECT review.*, user.username FROM review INNER JOIN user ON review.user_id = user.user_id WHERE review.movie_id = ?", Movie_id).Scan(&review).Count(count)
//	return count
//}
