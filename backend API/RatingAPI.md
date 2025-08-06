

# Rating

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content 
 rating | result class
 
## AddRating
Users can add rating for all the movies.  
**Method**:POST  
**URL**:http://localhost:8000/user/rating/addRating  
**Auth required**: Need token  
If add rating error   
```
{
  "message": "add rating error",
  "code":    404,
}
```  
If add rating success
```
{
 "message": "add rating success",
 "code":    200,
 "rating":   res,
}
```    
The result class will be same as the rating struct.

## DeleteRating
Users can delete rating for all the movies.  
**Method**:POST  
**URL**:http://localhost:8000/user/rating/deleteRating  
**Auth required**: Need token  
If delete rating error   
```
{
  "message": "delete rating error",
  "code":    404,
}
```  
If delete rating success
```
{
 "message": "delete rating success",
 "code":    200,
 "rating":   res,
}
```   
The result class will be same as the genre struct.

## UpdateRating
Users can update rating for all the movies.  
**Method**:POST   
**URL**:http://localhost:8000/user/rating/updateRating  
**Auth required**: Need token   
If update rating error   
```
{
  "message": "update rating error",
  "code":    404,
}
```  
If update rating success
```
{
 "message": "update rating success",
 "code":    200,
 "rating":   res,
}
```   
The result class will be same as the rating struct.

## ReadRating
users can read all the ratings for the movies they wrote before. 
**Method**:POST   
**URL**:http://localhost:8000/user/rating/readRating  
**Auth required**: Need token
If search rating error   
```
{
  "message": "search rating error",
  "code":    404,
}
```  
If search rating success
```
{
 "message": "search rating success",
 "code":    200,
 "rating":   res,
}
``` 
The result class will be same as the rating struct.

## ComputeAvgGrade
the average grade for each movie will be updated according to all the users' rating. 
**Method**:POST   
**URL**:http://localhost:8000/user/rating/computeAvgGrade  
**Auth required**: do not need token   
If search rating error   
```
{
  "message": "search rating error",
  "code":    404,
}
```  
If search rating success
```
{
 "message": "search rating success",
 "code":    200,
 "rating":   res,
}
``` 
The result class will be same as the rating struct.
