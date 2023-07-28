[Back](https://github.com/denitiawan/research-swagger-gomod-gin/blob/main/README.md)

## How to setup swagger on Go Project + Gorila Mux

![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/89fcd4f1-3f89-465f-b510-d589cc92e642)

###  Install Swagger-CLI

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

### Install Swagger librarries

```
github.com/swaggo/http-swagger/v2 v2.0.1
github.com/swaggo/swag v1.8.1
```

### Generate Swagger Docs
```
# run command on terminal
swag init
```
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/f215b733-1adf-4a9d-a903-ceea25b4dfee)


### Add Swagger Annotation on Main.go
```
// --------------------------------------------------
// annotations		: Annotation used for Swagger-UI
//					and will be mapping to folder and files (./root/docs/**)
// docs import		: import ( _ "denitiawan/research-swagger-gomod-gin/docs" )
//					will be used for update all values on all files inside that folder
//					when you run syntax (swag init)
// url swagger-ui 	: http://localhost:5050/nexsoft/doc/api/swagger/index.html
// --------------------------------------------------
// @version		1.1.0
// @title 		Demo Swagger-UI (GO+GORILLA MUX) for Nexsoft Project
// @description Implement swagger-ui on Go project with Gorilla Mux (web framework) + JWT Authorization
// @host 		localhost:5050
// @BasePath 	/

// ------showing authorize button (but validation jwt is not working)---------
// @Security Authorization
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @schemes http
// ------showing authorize button (but validation jwt is not working)---------
func main() {
```
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/3def6f72-87b8-4bc1-b714-2d75e9ea5e9a)


[Main.go](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/main.go)
### Add Swagger Annotation on Controller
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






### Add SwaggerRouting and call on main.go
- path swagger is customable
- path (/*any) is required for load the html page own by swagger
```

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

// path swagger is customable
// path (/*any) is required for load the html page own by swagger
// http://localhost:8810/nexsoft/doc/api/swagger/index.html
func SwaggerRouting(router *mux.Router) {
	prefix := "/nexsoft/doc/api"
	router.PathPrefix(prefix).Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		//httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		//httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

}

```
```
main.go{
...
...
config.SwaggerRouting(appRoute)
...
...
```
[SwaggerConfig.go](https://github.com/denitiawan/research-swagger-gomod-gorillamux/blob/main/config/SwaggerRouting.go)

### Run Application
update swagger docs
```
swag init
```
run app
```
go run ./main.go
```
open swagger-ui on browser
```
http://localhost:5050/nexsoft/doc/api/swagger/index.html
```
![image](https://github.com/denitiawan/research-swagger-gomod-gorillamux/assets/11941308/2a96faa5-5091-403a-a731-a4f6c3c8cb78)

