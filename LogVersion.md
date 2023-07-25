# Version Rule

1 [Major Version : Major changes & break the bussines rule & break the API ]
2 [Minor Version : Minor change & not break the bussines rule & add,edit,delete some API & Optimizing]
3 [Patches : Bugfixing ]

## version : 1.0.0 | release date : 19/07/2023
- first creating the project
- create controller (login, user, product)


## version : 1.1.0 | release date : 21/07/2023
- implement swagger-cli
- implement swagger annotation on main.go 
- implement swagger annotation on controller
- configure app info for swagger-ui
- configure controller for swagger-ui

## version : 1.2.0 | release date : 22/07/2023
- add migration database 

## version : 1.3.0 | release date : 25/07/2023
- implement CRUD funtion to UserRepo  (Create, Update , Delete, FindById, FindAll)

## version : 1.4.0 | release date : 25/07/2023
- Config : implement load Config using librarry (spf13/viper)
- Env : impelement .env for split dev and docker environtment
- Create User API : implement generate password (OAuth0) JWT
- login API : implement genereate Token JWT Bearer
- Middleware : implement validation JWT token when user hit all APIs

## version : 1.5.0 | release date : 25/07/2023
- implement swagger Authorization button (button authorize show, but cannot impact to all APIs when execute )

## version : 1.6.0 | release date : 25/07/2023
- swagger doc : add param type header "Authorization" for validate & secure all APIs when execute





