#Microservice used to search netflix movies and series

Build application
> go build microservice.go

Build docker image
> docker build -t netflix/microservice:1.0.0 .

Run docker container
> docker run --rm --network=host -p 8080:8080 netflix/microservice:1.0.0 postgres://netflix_user:TGxCOWiDbpcwgSM@localhost:5432/netflix?sslmode=disable

Run tests
> go test -v
> go test -cover
