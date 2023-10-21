package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)




type Handler struct{
	Router *mux.Router
	Service MovieService
	Server *http.Server
}


func NewHandler(service MovieService) *Handler{
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Server = &http.Server{
		Addr: ":8080",
		Handler: h.Router,
	}
	return h
}


func (h *Handler) mapRoutes(){
	h.Router.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World")
	}  )
	h.Router.HandleFunc("/api/v1/movie", h.CreateMovie).Methods("POST")
	h.Router.HandleFunc("/api/v1/movie/{id}", h.GetMovie).Methods("GET")
	h.Router.HandleFunc("/api/v1/movie/{id}", h.UpdateMovie).Methods("PUT")
	h.Router.HandleFunc("/api/v1/movie/{id}", h.DeleteMovie).Methods("DELETE")
}


func (h *Handler) Serve() error{
	go func ()  {
		if err := h.Server.ListenAndServe(); err != nil{
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("shut down gracefully")
	
	return nil
}