

# watch List

parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 watchList | result class
 
## AddWatchList
Users can add movies they have watched in the watched list.  
**Method**:POST  
**URL**:http://localhost:8000/user/watchedList/addWatchList  
**Auth required**: Need token  
If add watchList error   
```
{
  "message": "add to WatchList error",
  "code":    404,
}
```  
If add watched success
```
{
 "message": "add to watch success",
 "code":    200,
 "watchList":    res,
}
```    
The result class will be same as the watched_list struct.

## DeleteWatchList
Users can delete movies they have watched in the watched list. 
**Method**:POST  
**URL**:http://localhost:8000/user/watchedList/deleteWatchList  
**Auth required**: Need token  
If delete watched error   
```
{
  "message": "delete from WatchList error",
  "code":    404,
}
```  
If delete watched success
```
{
 "message": "delete from WatchList success",
 "code":    200,
 "watchList":    res,
}
```   
The result class will be same as the review struct.

## ReadWatchList
Users can see all the movies they have watched in the watched list. 
**Method**:POST   
**URL**:http://localhost:8000/user/watchedList/readWatchList  
**Auth required**: Need token   
If read watched list error   
```
{
  "message": "read watch list error",
  "code":    404,
}
```  
If read watched list success
```
{
 "message": "read watch list success",
 "code":    200,
 "watchList":    res,
}
```   
The result class will be same as the review struct.
