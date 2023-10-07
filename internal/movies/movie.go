package movies

import (
	"context"
	"fmt"
	"time"
)

type Movie struct {
	ID          string
	Title       string
	Slug        string
	Description string
	Producer    string
	Duration    time.Duration
}

type Store interface{
	GetMovie(context.Context, string)(Movie, error)
}

type Service struct{
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
	if err != nil{
		fmt.Println(err)
		return Movie{}, nil
	}

	return mv, nil
}
