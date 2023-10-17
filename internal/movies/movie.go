package movies

import (
	"context"
	"errors"
	"fmt"


)

var (
	ErrFetchingMovie = errors.New("failed to fetch movie by id")
)

type Movie struct {
	ID          string
	Title       string
	Slug        string
	Description string
	Producer    string
	Duration    string
	Author      string
}

type Store interface {
	GetMovie(context.Context, string) (Movie, error)
	CreateMovie(context.Context, Movie) (Movie, error)
	DeleteMovie(context.Context, string) error
	UpdateMovie(context.Context, string,  Movie) (Movie, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetMovie(ctx context.Context, id string) (Movie, error) {
	fmt.Println("retrieve movie")
	mv, err := s.Store.GetMovie(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Movie{}, ErrFetchingMovie
	}

	return mv, nil
}

func (s *Service) UpdateMovie(ctx context.Context,id string, updatedMv Movie) (Movie, error) {
	mv, err :=s.Store.UpdateMovie(ctx, id, updatedMv)
	if err != nil{
		fmt.Println("error updating movie")
		return Movie{}, err
	}
	return mv, nil
}

func (s *Service) DeleteMovie(ctx context.Context, id string) error {
	return s.Store.DeleteMovie(ctx, id)
}

func (s *Service) CreateMovie(ctx context.Context, mv Movie) (Movie, error) {
	insertedMv, err := s.Store.CreateMovie(ctx, mv)
	if err != nil {
		return Movie{}, err
	}
	return insertedMv, nil
}
