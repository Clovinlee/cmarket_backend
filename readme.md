
# Backend of CMarket using GO

This project is a microproject of CMarket, a mockup of ecommerce catalogue website using GoLang, Gin, and GORM.

#### Project highlight
- Rest API of Go
- Pagination
- Lazy Loading
- Advanced filter & queries

# Prerequisite
- go version of 1.22.2
- git
- .env

## .env
env consist of :
```
PORT=8000
DATABASE_URL="host={HOST} user={USER} password={PASS} dbname={DBNAME}"
```
As for the complete database, please contact me since I don't provide a database seed / migration. Or run yourself with raw_db.sql

# How to run
``` go
// getting the dependency
go mod download
// running the project
go run . 
// or use daemon tools package
CompileDaemon -command="./cmarket"
```