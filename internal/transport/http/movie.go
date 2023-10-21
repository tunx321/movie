package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tunx321/movie/internal/movies"
)


type MovieService interface{
	CreateMovie(context.Context, movies.Movie)(movies.Movie, error)
	GetMovie(ctx context.Context, id string) (movies.Movie, error)
	UpdateMovie(ctx context.Context, id string, newMv movies.Movie) (movies.Movie, error)
	DeleteMovie(ctx context.Context, id string) error
}


type Response struct{
	Message string
}

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request){
	var mv movies.Movie
	if err := json.NewDecoder(r.Body).Decode(&mv); err != nil{
		return 
	}

	mv, err := h.Service.CreateMovie(r.Context(), mv)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(mv); err != nil{
		panic(err)
	}
}

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mv, err := h.Service.GetMovie(r.Context(), id)
	if err != nil{
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(mv); err != nil{
		panic(err)
	}
}

func (h *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var mv movies.Movie
	if err := json.NewDecoder(r.Body).Decode(&mv); err != nil{
		return
	}


	mv, err := h.Service.UpdateMovie(r.Context(), id, mv)

	if err != nil{
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(mv); err != nil{
		panic(err)
	}
}

func (h *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteMovie(r.Context(), id)
	if err != nil{
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil{
		log.Print(err)
		panic(err)
	}

}	
