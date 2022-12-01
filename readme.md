### Carlord Backend

#### Requirement:
##### `docker` or `go 1.19.3`+`postgres`

#### Run:
    docker-compose up

or 
    
    go mod download
    go build
    
then run the binary.

default port: 8686

#### swagger
swagger needs documentations on api function to work.
to update swagger
    
    go mod download
    swag init

visit :8686/api/

