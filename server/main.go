package main

import (
	"bytes"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"./postgresql"
)

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	FatherName  string  `json:"fathername"`
	PhoneNumber string  `json:"phonenumber"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Birthday    string  `json:"birthday"`
	Rating      float64 `json:"rating"`
}

type Order struct {
	ID             int       `json:"id"`
	BookExemplarID []int     `json:"bookexemplarid"`
	ClientID       int       `json:"clientid"`
	OrderDate      time.Time `json:"orderdate"`
	RefundDate     time.Time `json:"refunddate"`
}

type BookExemplar struct {
	ID     int `json:"id"`
	BookID int `json:"bookid"`
	State  int `json:"stateid"`
}

type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PublisherID int       `json:"publisherid"`
	GenreID     int       `json:"genreid"`
	BookTypeID  int       `json:"booktypeid"`
	BooksQty    int       `json:"booksqty"`
	WritingDate time.Time `json:"writingdate"`
}

type AuthorBook struct {
	AuthorID int `json:"authorid"`
	BookID   int `json:"bookid"`
}

type Authors struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	FatherName string    `json:"fathername"`
	Birthday   time.Time `json:"birthday"`
}

type Publishers struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Postcode    string `json:"postcode"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BookType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	RoomID    int       `json:"roomid"`
	UserID    int       `json:"clientid"`
	EventDate time.Time `json:"eventdate"`
	PeopleQty int       `json:"peopleqty"`
}

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

// Users - хранилище для отправки ответов сервера фронтенду
var Users = map[int]User{
	// "1": {
	// 	ID:          1,
	// 	Name:        "Bob",
	// 	Surname:     "Popov",
	// 	FatherName:  "Popovich",
	// 	PhoneNumber: "88005553535",
	// 	Email:       "bob-popov@itmo.ru",
	// 	Password:    "bibibaba1030",
	// 	Birthday:    time.Date(2000, 4, 17, 0, 0, 0, 0, time.UTC).Format("02.01.2006"),
	// 	Rating:      5.0,
	// },
	// //.... здесь другие клиенты
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {

	// Проверяем POST-запрос или нет
	if r.Method == http.MethodPost {
		var user User
		var buf bytes.Buffer
		// читаем тело запроса
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// десериализуем JSON в Artist
		if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Users[user.ID] = user
	}
	resp, err := json.Marshal(Users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	Users = postgresql.SelectAllUsers()
	http.HandleFunc(`/`, JSONHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}
