package main

import (
	"fmt"
	"net/http"

	h "library_project/server/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/library/users", h.GetUsers)
	r.Get("/library/veruser", h.VerifyUser)
	r.Get("/library/orders", h.GetOrders)
	r.Get("/library/bookex", h.GetBookEx)
	r.Get("/library/book", h.GetBook)
	r.Get("/library/bookauthor", h.GetAuthorBook)
	r.Get("/library/author", h.GetAuthor)
	r.Get("/library/publisher", h.GetPublisher)
	r.Get("/library/genre", h.GetGenre)
	r.Get("/library/series", h.GetSeries)
	r.Get("/library/event", h.GetEvent)
	r.Get("/library/room", h.GetRoom)
	r.Get("/library/bindings", h.GetBinding)
	r.Delete("/library/event", h.DeleteEvent)
	r.Post("/library/event", h.PostEvent)
	r.Post("/library/orders", h.PostOrder)
	r.Post("/library/users", h.PostUser)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}
