## Movie echo grpc

```sh
migrate -path ./migrations -database "postgres://holyraven:holyraven@localhost:5432/crudsqlc?sslmode=disable" up
```

### Curl

```sh
curl -X GET http://localhost:5000/movies

curl -X GET http://localhost:5000/movies/1

curl -X POST -H "Content-Type: application/json" -d '{"Title": "New Movie", "Genre": "Action"}' http://localhost:5000/movies

curl -X PUT -H "Content-Type: application/json" -d '{"Title": "Updated Movie", "Genre": "Comedy"}' http://localhost:5000/movies/1

curl -X DELETE http://localhost:5000/movies/1
```
