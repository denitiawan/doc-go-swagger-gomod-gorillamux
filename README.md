## About
Research how to implement swagger on go (go.mod) using gorilla mux & net/http
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/629ea8be-b404-4a55-bae8-f5846b3f6930)

## Librarries

```
# go version
go 1.18

# Web Framework & Router
github.com/gorilla/mux
net/http

# Validator DTO
github.com/go-playground/validator/v10 v10.14.1

# Log
github.com/rs/zerolog v1.29.1

# API Docs
go install github.com/swaggo/swag/cmd/swag@v1.8.12
github.com/swaggo/http-swagger/v2 v2.0.1
github.com/swaggo/swag v1.8.1
  

 v1.8.12 

# Database & Driver
database/sql

# Database Migrate
github.com/rubenv/sql-migrate
go install github.com/rubenv/sql-migrate/...@latest
go mod download github.com/gobuffalo/packr/v2
go mod tidy

# Database Server
Mysql
```

## Create mysql container
[docker-compose.yml](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/docker-compose.yml)
```
# run syntax
docker-compose up -d

# docker will create docker container (mysql)
host: localhost
port: 3332      
user: user
password: password
database : db_mysql
```
## Run App
- open terminal
```
- go mod tidy
- go get
- go run ./main.go
```

## postman
[postman](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/Swagger-GO.18-v%201.0.0.postman_collection.json)



## How to setup swagger on (go.mod & Gorilla Mux)
- [Setup Swagger in GO + Gorilla Mux](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/tmp/readme/setup_swagger.md)
