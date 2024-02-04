package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	pb "movie-grpc-echo/proto"
	"movie-grpc-echo/repository"
	db "movie-grpc-echo/schema"
	"movie-grpc-echo/service"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var DB *db.Queries

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func init() {
	DatabaseConnection()
}

func DatabaseConnection() *db.Queries {
	host := "localhost"
	port := "5432"
	dbName := "crudsqlc"
	dbUser := "holyraven"
	password := "holyraven"

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	DB = db.New(conn)

	err = conn.Ping()

	if err != nil {
		log.Fatal("Error pinging the database...", err)
	}

	fmt.Println("Database connection successful...")

	return DB
}

func main() {
	fmt.Println("gRPC server running ...")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	context := context.Background()

	movieRepo := repository.NewMovieRepository(DB, context)

	movieService := service.NewMovieService(movieRepo)

	s := grpc.NewServer()

	pb.RegisterMovieServiceServer(s, movieService)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
