package postgresql

import (
	"database/sql"
	"fmt"

	// здесь тоже бы импортировать юзера из мейна, а не просто объявлять заново
	_ "github.com/lib/pq"
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

// SelectAllUsers возвращает мапу из базы данных, которая заполняется записями о пользователях из неё
func SelectAllUsers() map[int]User {

	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM postgres.public.user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := map[int]User{}

	for rows.Next() {
		u := User{}

		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.FatherName, &u.PhoneNumber, &u.Email, &u.Birthday, &u.Rating)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users[u.ID] = u
	}

	for i := 1; i <= len(users); i++ {
		fmt.Println(users[i])
	}
	return users
}

// AddUser добавляет запись в базу данных из структуры в параметре в таблицу пользователей
func AddUser(u *User) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//Adding new data
	testUser := User{ID: 4, Name: "testName", Surname: "tSur", FatherName: "tFat", PhoneNumber: "8990001003040", Email: "bibbib@riu", Birthday: "10.03.2007", Rating: 3.0}
	_, err = db.Exec("INSERT INTO public.user (name, surname, father_name, phone_number, email, birthday, rating) VALUES ($1, $2, $3, $4, $5, $6, $7)", testUser.Name, testUser.Surname, testUser.FatherName, testUser.PhoneNumber, testUser.Email, testUser.Birthday, testUser.Rating)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("rows were appended successfully")
	}

	defer db.Close()
}
