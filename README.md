# IMDB Clone System

A full-stack movie database application inspired by IMDB, built with Vue.js frontend and Go backend.

## 🎯 Features

- **User Authentication**: Secure JWT-based authentication system
- **Movie Database**: Browse, search, and discover movies
- **Ratings & Reviews**: Rate movies and write reviews
- **Actor Information**: Search for actors and their filmographies
- **Watchlist**: Save movies to your personal watchlist
- **Admin Panel**: Admin functionality for content management

## 👥 Team Members

### Frontend Development
- Yuxuan Wang (7566-9009)
- Shihuan Wang (9715-8829)

### Backend Development
- Tao Zhang (7636-6624)  
- Zihan Guo (8615-3487)

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.19+
- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: MySQL 8.0+
- **Authentication**: JWT (JSON Web Tokens)
- **CORS**: Custom middleware for cross-origin requests

### Frontend  
- **Language**: JavaScript (ES6+)
- **Framework**: Vue.js 2.7
- **UI Library**: Element UI
- **HTTP Client**: Axios
- **Date Handling**: Day.js
- **Routing**: Vue Router
- **Build Tool**: Webpack

## 📋 Prerequisites

- **Node.js**: v16.0+ and npm
- **Go**: v1.19+
- **MySQL**: v8.0+
- **Git**: Latest version

## ⚙️ Installation & Setup

### 1. Clone the Repository
```bash
git clone <repository-url>
cd IMDB_Clone-main
```

### 2. Environment Configuration
Create a `.env` file in the root directory:
```bash
cp .env.example .env
```
Update the `.env` file with your configuration:
```env
# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=IMDB
DB_HOST=localhost
DB_PORT=3306
DB_CHARSET=utf8
DB_DEBUG=false

# JWT Configuration
JWT_SECRET_KEY=your_secret_jwt_key

# Server Configuration
SERVER_PORT=8000
```

### 3. Database Setup
1. Install and start MySQL
2. Create the database:
```sql
CREATE DATABASE IMDB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
3. Import the database schema:
```bash
mysql -u your_username -p IMDB < IMDB.sql
```

### 4. Backend Setup
```bash
cd backend
go mod tidy
go run main.go
```
The backend server will start on `http://localhost:8000`

### 5. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
The frontend application will be available at `http://localhost:8080`

## 🚀 Running the Application

### Development Mode
1. Start the backend server:
```bash
cd backend && go run main.go
```

2. Start the frontend development server:
```bash
cd frontend && npm run dev
```

### Production Build
```bash
cd frontend && npm run build
```

## 🎯 API Endpoints

### Authentication
- `POST /user/logIn` - User login
- `POST /user/signUp` - User registration
- `POST /user/checkUsername` - Check username availability

### Movies
- `POST /user/movie/searchMovieByName` - Search movies by name
- `POST /user/movie/searchMovieByCast` - Search movies by cast
- `POST /user/movie/searchMovieByMovieId` - Get movie by ID
- `POST /user/movie/topMovie` - Get top-rated movies

### Reviews & Ratings
- `POST /user/review/addReview` - Add movie review
- `POST /user/rating/addRating` - Rate a movie
- `POST /user/rating/computeAvgGrade` - Get average rating

### Watchlist
- `POST /user/watchedList/addWatchedList` - Add to watchlist
- `POST /user/watchedList/readWatchedList` - Get user's watchlist

## 🧪 Testing

### Backend Tests
```bash
cd backend && go test ./...
```

### Frontend Tests
```bash
cd frontend && npm run test
```

### E2E Tests
```bash
cd frontend && npm run e2e
```

## 📈 Recent Improvements

### Security Enhancements
- **Environment Variables**: Database credentials and JWT secrets now use environment variables
- **JWT Security**: Updated to latest JWT library with improved security
- **Password Protection**: Passwords are no longer returned in API responses
- **CORS Middleware**: Added proper CORS handling for cross-origin requests

### Code Quality
- **Unified Error Handling**: Consistent API response format across all endpoints
- **Dependency Updates**: Updated both Go and npm dependencies to latest stable versions
- **Database Optimization**: Improved connection pooling and query optimization

### Performance Improvements  
- **Frontend Optimizations**: Updated Vue.js and related packages for better performance
- **Date Handling**: Replaced moment.js with lighter day.js library
- **Connection Management**: Optimized database connection settings

## 🗺️ Project Architecture
<img src="/IMDB.png" alt="IMDB Clone System Architecture"/>

## 📝 Development Notes

For detailed information about each sprint and development milestones, please refer to:
- [Sprint 1](Sprint%201.md)
- [Sprint 2](Sprint%202.md)
- [Sprint 3](Sprint%203.md)
- [Sprint 4](Sprint%204.md)
