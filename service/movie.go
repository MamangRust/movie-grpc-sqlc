package service

import (
	"context"
	"movie-grpc-echo/domain"
	pb "movie-grpc-echo/proto"
	"movie-grpc-echo/repository"
)

type movieService struct {
	repo repository.MovieRepository
	pb.UnimplementedMovieServiceServer
}

func NewMovieService(repo repository.MovieRepository) *movieService {
	return &movieService{
		repo: repo,
	}
}

func (s *movieService) CreateMovie(ctx context.Context, request *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	movie := domain.Movie{
		Title: request.Movie.Title,
		Genre: request.Movie.Genre,
	}

	err := s.repo.CreateMovie(movie)

	if err != nil {
		return nil, err
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil

}

func (s *movieService) GetMovies(ctx context.Context, request *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	movies, err := s.repo.GetMovies()

	if err != nil {
		return nil, err
	}

	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil

}

func (s *movieService) GetMovie(ctx context.Context, request *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	movie, err := s.repo.GetMovie(request.Id)

	if err != nil {
		return nil, err
	}

	return &pb.ReadMovieResponse{
		Movie: movie,
	}, nil
}

func (s *movieService) UpdateMovie(ctx context.Context, request *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	movie := domain.Movie{
		ID:    int(request.Movie.Id),
		Title: request.Movie.Title,
		Genre: request.Movie.Genre,
	}

	res, err := s.repo.UpdateMovie(movie)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateMovieResponse{
		Movie: res,
	}, nil

}

func (s *movieService) DeleteMovie(ctx context.Context, request *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	err := s.repo.DeleteMovie(request.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
