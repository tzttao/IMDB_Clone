# Description of IMDB
IMDB_Clone is a web application that we mimic the IMDB website. User could browse and search movie and cast information and detail. Meanwhile, they can write a comment or rate the movie they like, but prerequisite is they need to login. We also provide a sign up function for users who like our website. For users who already login, they can update or delete their account. For Administrator, they can conduct the movie and cast data, such as add, update or delete data. For some inappropriate comments, admin could delete the commemt, even delete user account.

[Demo video functionality](https://youtu.be/X4Ci1AL6kNk)

[Cypress Test video](https://youtu.be/JgWZIkhgBe4)

[Backend unit test video](https://youtu.be/ZRYLe-DpBHs)  
We saw the annoucement that do not link an outside webside, but these videos are more than 100M, github cannot allow us to upload. Sorry about that.

[Link to API Documentation](https://github.com/BazingaOo/IMDB_Clone/tree/main/backend%20API)  

[Link to Project board](https://github.com/BazingaOo/IMDB_Clone/projects)

[Link to Sprint4 deliverables](https://github.com/BazingaOo/IMDB_Clone.git) Note: The whole project is in the frontend and backend folders.
# What we accomplished:
## Front-End

Abstract Summary

created the actor information page\
created the movie information page\
created the comment page\
implemented the comment function\
implemented the rating function\
implemented the ralative movie function\
embellished all the pages previously created\
redesigned some pages\
cypress test for login\
cypress test for searching actors\
cypress test for searching movies


### Adding more Details in HomePage
For aesthetic design of our homepage, we refill the picture size to 
make movie posters showing like IMax screen. \
Moreover, we find more suitable-size pictures as our Top-Movies' covers.

### Implementing new Functionalities
#### For Searching Movie Functionality
Add the grade of movies\
Add the description of actors\
Complete the searching function with added movies and actors in mysql database

#### For Movie Information Page
Show the genre, cast, abstract, and user's rating of searching movie\
Enable the user leave their rating on searching movie and calculate the average score\
Add the comment functionality to let users leave their comments on searching movies

#### For Leaving Comments Page
Implement the type-in box for users to leave some comments on movies\
Update the submitted comments in database and show comments on Movie Info page

#### For Actor Information Page
Add description of the searched actor\
Add the movie which is related to the searched actor and Add jump to the related movie info page



## Back-End
### For Comment model
basic CURD for comment    
add a pagination to display comments for each movie  
add jwt middleware to the router  
### For Rating model
basic CURD for rating    
add a function that could display the trending movie  
add a computing average socore function for each function  
add jwt middleware to the router  
### For Genre model
basic CURD for genre  
add genre to the search result for each movie  
add jwt middleware to the router  
# How to run
## Install npm package
```
npm install
```
## Run the front server
```
npm run serve
```
the homepage is localhost:8080
## Run the backend server
open in the GoLand software and config the go build, then click the start button, the backend address is localhost:8000

# Frontend and backend team members
## Group Members
### front-end
Yuxuan Wang, 7566-9009<br/>
Shihuan Wang, 9715-8829<br/>
### back-end
Tao Zhang, 7636-6624<br/>
Zihan Guo, 8615-3487<br/>
