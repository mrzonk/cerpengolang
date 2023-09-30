# Go Cerpen Rest Api
A RESTful API example for simple Short Story with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run
```bash
# Download this project
go get https://github.com/mrzonk/cerpengolang.git
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/mrzonk/cerpengolang/blob/master/config/database.go)
```go
	username := ""
	password := ""
	host := ""
	port := ""
	database := ""
	connection := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

```

```bash
# Build and Run
cd cerpengolang
go build
./cerpengolang

# API Endpoint : http://127.0.0.1:3434/api
```



## API

#### /cerpens
* `GET` : Get all Cerpen
* `POST` : Create a new Cerpen

#### /cerpens/:id/details
* `GET` : Get a detail Cerpen

#### /cerpens/:id/update
* `PUT` : Update a Cerpen

#### /cerpens/:id/delete
* `Delete` : Delete One Project

#### /comment
* `GET` : Get all Comment
* `POST` : Create a new Comment

#### /cerpens/:id/details
* `GET` : Get a detail Comment

#### /cerpens/:id/update
* `PUT` : Update a Comment

#### /cerpens/:id/delete
* `Delete` : Delete One Comment

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 