[Back](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/README.md)

## How to setup swagger on Go Project + Gorila Mux
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/bb575a1f-3fd1-4a0b-a253-bfb50815abc0)

### 1. Install Swagger-CLI

install swag cli

```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@v1.8.12
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
// --[swagger notes]------------------------------------------
// annotations		: Annotation used for Swagger-UI
//					and will be mapping to folder and files (./root/docs/**)
// docs import		: import ( _ "denitiawan/research-swagger-gomod-gin/docs" )
//					will be used for update all values on all files inside that folder
//					when you run syntax (swag init)
// url swagger-ui 	: http://localhost:8899/nexsoft/doc/api/swagger/index.html

// --[swagger annotation]--------------------------------
// @version		1.1.0
// @title 		Demo Swagger-UI (GO+GIN) for Nexsoft Project
// @description Implement swagger-ui on Go project with gin (web framework) + JWT Authorization
// @host 		localhost:8899
// @BasePath 	/api

// --[showing authorize button (but validation jwt is not working)]---------
// @Security Authorization
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @schemes http

func main() {
....
....
}
```
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/a8753ccb-6269-4b6c-bae6-899339fa07ff)

[main.go](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/main.go)

### 5. Add Swagger annotation on controller

add annotation above all controller function (CRUD)

```
// @Tags			user
// @Router			/v1/user/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body user.UserDto true "UserDto.go"
// @Param			Authorization header string true "Authorization"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Create(ctx *gin.Context) {
	...
	...
	...
}

```
![image](https://github.com/denitiawan/research-swagger-gomod-gin/assets/11941308/8e613ca2-e6a8-4767-bd02-8ef7d65687ca)
[UserController.go](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/module/user/user_controller.go)
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
http://localhost:8810/nexsoft/doc/api/swagger/index.html
```
