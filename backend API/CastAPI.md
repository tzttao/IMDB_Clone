# Casts

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 cast | result class
 
## AddCast
Administrators can add new cast in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/addCast  
**Auth required**: Need token  
If add cast error   
```
{
  "message": "add cast error",
  "code":    404,
}
```  
If add cast success
```
{
 "message": "add cast success",
 "code":    200,
 "cast":   res,
}
```    
The result class will be same as the cast struct.

## DeleteCast
Administrators can delete cast in the system.  
**Method**:POST  
**URL**:http://localhost:8000/admin/DeleteCast  
**Auth required**: Need token  
If delete cast error   
```
{
  "message": "delete cast error",
  "code":    404,
}
```  
If delete cast success
```
{
 "message": "delete casts success",
 "code":    200,
 "cast":   res,
}
```   
The result class will be same as the cast struct.

## UpdateCast
Administrators can update Cast in the system.  
**Method**:POST   
**URL**:http://localhost:8000/admin/updateCast  
**Auth required**: Need token   
If update Cast error   
```
{
  "message": "update Cast error",
  "code":    404,
}
```  
If update Cast success
```
{
 "message": "update Cast success",
 "code":    200,
 "cast":   res,
}
```   
The result class will be same as the cast struct.

## SearchCastByName
search cast by a specific cast name.  
**Method**:POST   
**URL**:http://localhost:8000/searchCast  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the Cast",
  "code":    404,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "cast":   res,
}
``` 
The result class will be same as the cast struct.

## SearchCastById
search cast by a specific cast id.  
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchCastById  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the Cast",
  "code":    404,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "cast":   res,
}
``` 
The result class will be same as the cast struct.


## SearchCastByMovieId
search cast by a specific movie id.  
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchCastByMovieId  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the Cast",
  "code":    404,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "cast":   res,
}
``` 
The result class will be same as the cast struct.

## SearchRelativeMovieByCastId
search relative movie by a specific cast id.  
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchCastByMovieId  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "movie":   res,
}
``` 
The result class will be same as the movie struct.

## SearchCastByName
search cast by a specific cast name.  
**Method**:POST   
**URL**:http://localhost:8000/user/cast/searchCastByName  
**Auth required**: do not need token   
If search cast error   
```
{
  "message": "not found the movie",
  "code":    404,
}
```  
If search cast success
```
{
 "message": "search cast success",
 "code":    200,
 "cast":   res,
}
``` 
The result class will be same as the cast struct.
