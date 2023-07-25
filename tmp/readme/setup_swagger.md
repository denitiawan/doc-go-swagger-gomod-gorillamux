[Back](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/README.md)

## How to setup swagger on (go.mod & GIN)
![image](https://github.com/denitiawan/research-swagger-gomod-gin/assets/11941308/961c63b7-eb50-42f7-9f54-3381631fd641)

### 1. Install Swagger-CLI

install swag cli

```
go install github.com/swaggo/swag/cmd/swag@v1.8.12
go get -u github.com/swaggo/swag/cmd/swag
```

test swag cli success installed

```
swag -h
```

```
NAME:
   swag.exe - Automatically generate RESTful API documentation with Swagger 2.0 for Go.

USAGE:
   swag.exe [global options] command [command options] [arguments...]

VERSION:
   v1.8.12

COMMANDS:
   init, i  Create docs.go
   fmt, f   format swag comments
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### 2. Install Swagger librarries

```
github.com/swaggo/files v1.0.1
github.com/swaggo/gin-swagger v1.6.0
```

### 3. Add Swagger route on main.go
- path swagger is customable
- path (/*any) is required for load the html page own by swagger
```
import(
    swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	)

...
...
router.GET("nexsoft/doc/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
...
...
```

### 4. Add Swagger annotation on main.go

do initial swagger doc

```
swag init
```

this comman will generate folder and files, and will update all swagger files when running (swag init)

```
./root
    /docs
        docs.go
        swagger.json
        swagger.yaml       
```

add import docs folder

```
_ "denitiawan/research-swagger-gomod-gin/docs"
```

add annotation above on main function

```
// @version		1.0
// @title 		Nexsoft Demo Swagger in GO
// @description How to implement swagger-ui in Go and gin http client project
// @host 		localhost:8899
// @BasePath 	/api
func main() {
....
....
}
```

swagger notes

```
// annotations		: Annotation used for Swagger-UI
//					and will be mapping to folder and files (./root/docs/**)
// docs import		: import ( _ "denitiawan/research-swagger-gomod-gin/docs" )
//					will be used for update all values on all files inside that folder
//					when you run syntax (swag init)
// url swagger-ui 	: http://localhost:8899/nexsoft/doc/api/swagger/index.html
```

### 5. Add Swagger annotation on controller

add annotation above all controller function (CRUD)

```
// @Tags			user
// @Router			/v1/user/save [post]
// @Summary			save
// @Description	                save
// @Param			RequestBody body user.UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Create(ctx *gin.Context) {
	...
	...
	...
}

```

### 6. open swagger ui on web
update swagger files
```
swag init
```
run app
```
go run ./main.go
```
open swagger-ui on browser
```
http://localhost:8899/nexsoft/doc/api/swagger/index.html
```
