# Movies

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 movie | result class
 
## AddNewMovies
Administrators can add new movies in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/movie/addMovie  
**Auth required**: Need token  
If add movies error   
```
{
  "message": "add movies error",
  "code":    404,
}
```  
If add movies success
```
{
 "message": "add movies success",
 "code":    200,
 "movie":   res,
}
```    
The result class will be same as the movie struct.

## DeleteMovies
Administrators can delete movies in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/movie/deleteMovie  
**Auth required**: Need token  
If delete movies error   
```
{
  "message": "delete movies error",
  "code":    404,
}
```  
If delete movies success
```
{
 "message": "delete movies success",
 "code":    200,
 "movie":   res,
}
```   
The result class will be same as the movie struct.

## UpdateMovies
Administrators can update movies in the system.  
**Method**:POST   
**URL**:http://localhost:8000/admin/movie/updateMovie  
**Auth required**: Need token   
If update movies error   
```
{
  "message": "update movies error",
  "code":    404,
}
```  
If update movies success
```
{
 "message": "update movies success",
 "code":    200,
 "movie":   res,
}
```   

## SearchMovies
All the users can search movies by name and a specific cast.
### SearchMoviesByName
All the users can search movies by name.  
**Method**:POST   
**URL**:http://localhost:8000/user/movie/searchMovieByName  
**Auth required**: do not need token   
If search movies error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search movies success
```
{
 "message": "search movies success",
 "code":    200,
 "movie":   res,
}
``` 
The result class will be same as the movie struct.

## SearchMovieByMovieId
search movies by a specific moive id.  
**Method**:POST   
**URL**:http://localhost:8000/user/movie/searchMovieByMovieId  
**Auth required**: do not need token   
If search movies error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search movies success
```
{
 "message": "search movies success",
 "code":    200,
 "result":   res,
}
``` 
Result Class
parameter  | discription
 ---- | ----- 
 MovieID  | movie id 
 MovieName  | movie name 
 MovieYear | the year of the movie
 castName | the name of the cast that user search
 castID | the id of the cast that user search
 
## SearchMoviesByCast
search movies by a specific cast.  
**Method**:POST   
**URL**:http://localhost:8000/user/movie/searchMovieByCast  
**Auth required**: do not need token   
If search movies error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search movies success
```
{
 "message": "search movies success",
 "code":    200,
 "result":   res,
}
``` 
Result Class
parameter  | discription
 ---- | ----- 
 MovieID  | movie id 
 MovieName  | movie name 
 MovieYear | the year of the movie
 castName | the name of the cast that user search
 castID | the id of the cast that user search
 
  
## SearchMovieByYear
search movies by year.  
**Method**:POST   
**URL**:http://localhost:8000/user/movie/searchMovieByYear  
**Auth required**: do not need token   
If search movies error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search movies success
```
{
 "message": "search movies success",
 "code":    200,
 "result":   res,
}
``` 
Result Class
parameter  | discription
 ---- | ----- 
 MovieID  | movie id 
 MovieName  | movie name 
 MovieYear | the year of the movie
 castName | the name of the cast that user search
 castID | the id of the cast that user search
 
## TopMovie
Search top 30 movies.
### SearchMoviesByName
All the users can search movies by name.  
**Method**:POST   
**URL**:http://localhost:8000/user/movie/topMovie  
**Auth required**: do not need token   
If search movies error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search movies success
```
{
 "message": "search movies success",
 "code":    200,
 "movie":   res,
}
``` 
The result class will be same as the movie struct.
