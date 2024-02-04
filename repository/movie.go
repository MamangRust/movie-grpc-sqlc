package repository

import (
	"context"
	"movie-grpc-echo/domain"
	pb "movie-grpc-echo/proto"
	db "movie-grpc-echo/schema"
)

type MovieRepository interface {
	CreateMovie(movie domain.Movie) error
	GetMovie(id int32) (*pb.Movie, error)
	GetMovies() ([]*pb.Movie, error)
	UpdateMovie(movie domain.Movie) (*pb.Movie, error)
	DeleteMovie(id int32) error
}

type movieRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewMovieRepository(db *db.Queries, ctx context.Context) *movieRepository {
	return &movieRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *movieRepository) GetMovies() ([]*pb.Movie, error) {
	res, err := r.db.GetMovies(r.ctx)

	if err != nil {
		return nil, err
	}

	pbMovies := make([]*pb.Movie, len(res))

	for i, m := range res {
		pbMovies[i] = &pb.Movie{
			Id:    m.ID,
			Title: m.Title,
			Genre: m.Genre,
		}
	}

	return pbMovies, nil
}

func (r *movieRepository) GetMovie(id int32) (*pb.Movie, error) {
	res, err := r.db.GetMovie(r.ctx, id)

	if err != nil {
		return nil, err
	}

	return &pb.Movie{
		Id:    res.ID,
		Title: res.Title,
		Genre: res.Genre,
	}, nil
}

func (r *movieRepository) CreateMovie(request domain.Movie) error {

	data := db.CreateMovieParams{
		Title: request.Title,
		Genre: request.Genre,
	}

	_, err := r.db.CreateMovie(r.ctx, data)

	if err != nil {
		return err
	}

	return nil
}

func (r *movieRepository) UpdateMovie(request domain.Movie) (*pb.Movie, error) {
	data := db.UpdateMovieParams{
		ID:    int32(request.ID),
		Title: request.Title,
		Genre: request.Genre,
	}

	res, err := r.db.UpdateMovie(r.ctx, data)

	if err != nil {
		return nil, err
	}

	return &pb.Movie{
		Id:    res.ID,
		Title: res.Title,
		Genre: res.Genre,
	}, nil
}

func (r *movieRepository) DeleteMovie(id int32) error {
	err := r.db.DeleteMovie(r.ctx, id)

	if err != nil {
		return err
	}

	return nil
}
