# go-awesome
Simple web application to test GoLang features

## TODO
- Update APIs structure
- Add configurable environments
- Add Docker build file
- Add CircleCI build file
- Add Sonar support
- TBD

## API Docs Generation with Swagger

### Install Swagger For Go
```
go get -u github.com/swaggo/swag/cmd/swag
```

### Generate/Update Swagger Generated Docs
Run from the top level directory of this project:
```
swag init
```
The service will be recompiled to incorporate the changes.

After APIs update please don't forget to commit your changes.

### Swagger Interactive Page
```
http://localhost:8080/swagger/index.html
```