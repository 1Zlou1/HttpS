package main

import (
	"HttpS/service/friend_servise"
	"HttpS/service/repository/map_repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {

	log.Println("start program")

	repository := map_repository.New()

	friendService := friend_servise.New(repository)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/create", friendService.CreateFriend)

	r.Post("/make_friends", friendService.MakeFriend)

	r.Get("/friends/{user_id}", friendService.GetFriends)

	r.Delete("/{user}", friendService.DeleteFriend)

	r.Put("/{user_id}", friendService.UpdateAge)

	http.ListenAndServe(":8080", r)

}
