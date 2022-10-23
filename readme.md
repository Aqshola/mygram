# [My Gram](https://mygram-production.up.railway.app/)
Simple rest api, build for submission <b>Digitalent - Hacktiv8 Scalable Web Service with GO LANG</b> <br/>


## Folder Structure 
- Config (Store Database config)
- Entity (Store Table Entity)
- DTO (Store JSON struct)
- Repository (Store file repository for each use case)
- Service (Store file for each use case or service)
- Handlers (Store file for controller and route)
- Helpers (Store file contain function helper)


## How to
- Clone repository
- Run ``` go run main.go```
- Open ```localhost:8080/swagger/index.html ``` for see documentation
- <b>NOTE</b> when input token for Authorization in swagger, make sure use this format `Bearer YOUR_TOKEN_HERE` since current swaggo still not implemented Bearer token


## Framework
- Web : <b>Gin Framework</b>
- Validator : <b>Go Validator</b>
- JWT: <b>Jwt-Go</b>
- Documentation: <b>Swaggo</b>

## What i learn
- Understanding how to build rest api using Go Lang
- Implementing Clean Architecture folder based on [Programmer Zaman Now Repository](https://github.com/khannedy/golang-clean-architecture)
- Learning how documenting API using Swagger

## What next?
* [ ] Add testing and mock testing


