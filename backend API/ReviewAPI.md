
# Review

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 review | result class
 
## AddNewReview
Users can add new reviews in any movies.  
**Method**:POST  
**URL**:http://localhost:8000/user/review/addReview  
**Auth required**: Need token  
If add reviews error   
```
{
  "message": "add reviews error",
  "code":    404,
}
```  
If add reviews success
```
{
 "message": "add reviews success",
 "code":    200,
 "review":  res,
}
```    
The result class will be same as the review struct.

## DeleteReviews
Users can delete reviews in the system.  
**Method**:POST  
**URL**:http://localhost:8000/user/review/deleteReview  
**Auth required**: Need token  
If delete reviews error   
```
{
  "message": "delete reviews error",
  "code":    404,
}
```  
If delete reviews success
```
{
 "message": "delete reviews success",
 "code":    200,
 "review":  res,
}
```   
The result class will be same as the review struct.

## UpdateReviews
Users can update thier own reviews in the system.  
**Method**:POST   
**URL**:http://localhost:8000/user/review/updateReview  
**Auth required**: Need token   
If update reviews error   
```
{
  "message": "update reviews error",
  "code":    404,
}
```  
If update reviews success
```
{
 "message": "update reviews success",
 "code":    200,
 "review":  res,
}
```   
The result class will be same as the review struct.

## ReadReviewByUserId
Users can select thier own reviews in the system.  
**Method**:POST   
**URL**:http://localhost:8000/user/review/readReview  
**Auth required**: Need token   
If select reviews error   
```
{
  "message": "select reviews error",
  "code":    404,
}
```  
If select reviews success
```
{
 "message": "select reviews success",
 "code":    200,
 "review":  res,
}
```   
The result class will be same as the review struct.

## ReadReviewByMovieId
On the movie detailed page, user can see the other guys comments.  
**Method**:POST   
**URL**:http://localhost:8000/user/review/readReviewByMovieId  
**Auth required**: Need token   
If select reviews error   
```
{
  "message": "select reviews error",
  "code":    404,
}
```  
If select reviews success
```
{
 "message": "select reviews success",
 "code":    200,
 "review":  res,
}
```   
The result class will be shown as fllowing:

Result Class
parameter  | discription
 ---- | ----- 
 Review_id  | review id 
 User_id  | user id 
 MovieID  | movie id 
 Review_content  | the content of this review that the users wrote 
 Username | the username of users who wrote review for this movie before




