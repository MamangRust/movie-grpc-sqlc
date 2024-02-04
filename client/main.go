package main

import (
	"context"
	"flag"
	"net/http"
	"strconv"

	"movie-grpc-echo/domain"
	pb "movie-grpc-echo/proto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/movies", func(c echo.Context) error {
		res, err := client.GetMovies(context.Background(), &pb.ReadMoviesRequest{})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"movies": res.Movies})
	})

	e.GET("/movies/:id", func(c echo.Context) error {
		id := c.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		res, err := client.GetMovie(context.Background(), &pb.ReadMovieRequest{Id: int32(idInt)})
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"movie": res.Movie})
	})

	e.POST("/movies", func(c echo.Context) error {
		var movie domain.Movie

		if err := c.Bind(&movie); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		data := &pb.Movie{
			Title: movie.Title,
			Genre: movie.Genre,
		}
		res, err := client.CreateMovie(context.Background(), &pb.CreateMovieRequest{Movie: data})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{"movie": res.Movie})
	})

	e.PUT("/movies/:id", func(c echo.Context) error {
		id := c.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		var movie domain.Movie

		movie.ID = idInt

		if err := c.Bind(&movie); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		res, err := client.UpdateMovie(context.Background(), &pb.UpdateMovieRequest{
			Movie: &pb.Movie{
				Id:    int32(movie.ID),
				Title: movie.Title,
				Genre: movie.Genre,
			},
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"movie": res.Movie})
	})

	e.DELETE("/movies/:id", func(c echo.Context) error {
		id := c.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		res, err := client.DeleteMovie(context.Background(), &pb.DeleteMovieRequest{Id: int32(idInt)})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		if res.Success == true {
			return c.JSON(http.StatusOK, map[string]interface{}{"message": "Movie deleted successfully"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error deleting movie"})
		}
	})

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
