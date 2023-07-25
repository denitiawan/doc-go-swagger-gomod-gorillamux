## About
Research how to implement swagger on go (go.mod) using gorilla mux & net/http


## Librarries
```
# Web Framework & Router
github.com/gorilla/mux
net/http

# Validator DTO
github.com/go-playground/validator/v10 v10.14.1

# Log
github.com/rs/zerolog v1.29.1

# API Docs
go install github.com/swaggo/swag/cmd/swag@v1.8.12
github.com/swaggo/http-swagger/example/gorilla/docs
github.com/swaggo/http-swagger/v2
  

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

## Run APP
[docker-compose.yml](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/docker-compose.yml)
```
# run syntax
docker-compose up -d

# docker will create docker image (swg_go_mysql)
host: localhost
port: 3399      
user: user
password: password
database : database
```

## postman
[postman](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/Swagger-GO.18-v%201.0.0.postman_collection.json)



## How to setup swagger on (go.mod & Gorilla Mux)
### 
### 
### 
