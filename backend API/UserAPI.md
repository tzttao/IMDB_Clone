# User 


parameter  | discription
 ---- | ----- 
 msg  | hint message 
 code  | status code 
 token | token content
 user | user class
## User sign up    

user sign up function

**Method**:POST  

**URL**:http://localhost:8000/user/signUp  

**Auth required**: do not Need token  

Sign up error   
```
{
  "msg": "sign up error",
  "code":    -1,
}
```  
Sign up success
```
{
  "msg": "sign up success",
  "code":    200,
}
```    

## User log in

user log in function

**Method**:POST  

**URL**:http://localhost:8000/user/logIn  

**Auth required**: do not Need token 

Login success
```
{
  "status":  0,
  "msg":     "login success",
  "user":    user,
  "token":   token,
}
```  

Login fail
```
{
  "msg": "username or password is wrong",
  "status":  -1,
}
```  

## CheckUsername
when user input userne on sign up page, check whether username unique or not

**Method**:POST  

**URL**:http://localhost:8000/user/CheckUsername

**Auth required**: Need token 

username unique
```
{
"message": "success",
}
```

username not unique
```
{
"message": "error",
}
```
## UpdateUserProfile
when user on their own profile page, could update theirselves information

**Method**:POST  

**URL**:http://localhost:8000/user/UpdateUserProfile

**Auth required**: Need token 

update error
```
{
  "message": "update error",
  "user":    user,
}
```

upate success
```
{
    "message": "update success",
    "user":    user,
}
```

## Delete User
when user on their own profile page, could delete user

**Method** GET  

**URL**:http://localhost:8000/user/DeleteUser

**Auth required**: Need token 

delete error
```
{
   "message": "delete error",
   "userId":  userId,
}
```

delete success
```
{
   "message": "delete success",
   "userId":  userId,
}
```
## GenerateToken
**Auth required**: Need token  
generate success
```
{
  "code":    200,
  "message": "login successfully",
  "status":  0,
  "user":    user,
  "token":   token,
}
```
