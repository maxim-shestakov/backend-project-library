package main

import (
	"fmt"
	"net/http"

	l "backend-project-library/server/dbconn"
	h "backend-project-library/server/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	l.Db = l.Connection(l.Db)
	defer l.Db.Close()
	r := chi.NewRouter()
	//r.Get("/library/users", h.GetUsers)
	r.Get("/library/orders", h.GetOrders)
	r.Get("/library/bookex", h.GetBookEx)
	r.Get("/library/book", h.GetBook)
	r.Get("/library/books", h.GetBooks)
	r.Get("/library/bookauthor", h.GetAuthorBook)
	r.Get("/library/author", h.GetAuthor)
	r.Get("/library/publisher", h.GetPublisher)
	r.Get("/library/genre", h.GetGenre)
	r.Get("/library/series", h.GetSeries)
	r.Get("/library/event", h.GetEvents)
	r.Get("/library/room", h.GetRoom)
	r.Get("/library/bindings", h.GetBinding)
	r.Get("/library/basket", h.GetBasket)
	r.Delete("/library/events", h.DeleteEvent)
	r.Post("/library/token", h.ReturnToken)
	r.Post("/library/events", h.PostEvent)
	r.Post("/library/orders", h.PostOrder)
	r.Post("/library/users", h.PostUser)
	r.Post("/library/baskets", h.PostBasket)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}
