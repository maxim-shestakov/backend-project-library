package postgresql

import (
	"database/sql"
	"fmt"

	"library_project/server/structures"

	_ "github.com/lib/pq"
)

// SelectAllUsers возвращает мапу из базы данных, которая заполняется записями о пользователях из неё
func SelectAllUsers() map[int]structures.User {

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

	users := map[int]structures.User{}

	for rows.Next() {
		u := structures.User{}

		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.FatherName, &u.PhoneNumber, &u.Email, &u.Birthday, &u.Rating, &u.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users[u.ID] = u
	}
	return users
}

// AddUser добавляет запись в базу данных из структуры в параметре в таблицу пользователей
func AddUser(u *structures.User) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//Adding new data

	_, err = db.Exec("INSERT INTO public.user (name, surname, father_name, phone_number, email, birthday, rating, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", u.Name, u.Surname, u.FatherName, u.PhoneNumber, u.Email, u.Birthday, u.Rating, u.Password)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("user was appended successfully")
	}

	defer db.Close()
}

func SelectUserData(u *structures.UserVer) structures.User {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.user WHERE password='%s' AND email='%s'", u.Password, u.Email)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	us := structures.User{}
	for rows.Next() {
		err := rows.Scan(&us.ID, &us.Name, &us.Surname, &us.FatherName, &us.PhoneNumber, &us.Email, &us.Birthday, &us.Rating, &us.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return us
}

func SelectAllOrders(u *structures.User) map[int]structures.Order {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.orders WHERE user_id=%d", u.ID)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Orders := map[int]structures.Order{}

	for rows.Next() {
		o := structures.Order{}

		err := rows.Scan(&o.ID, &o.BookExemplarID, &o.UserID, &o.OrderDate, &o.RefundDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Orders[o.ID] = o
	}
	return Orders
}

func SelectBookEx(id int) structures.BookExemplar {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.bookexemplar WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Bx := structures.BookExemplar{}

	for rows.Next() {
		err := rows.Scan(&Bx.ID, &Bx.BookID, &Bx.State)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return Bx
}

func SelectBook(id int) structures.Book {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.book WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	B := structures.Book{}

	for rows.Next() {
		err := rows.Scan(&B.ID, &B.Name, &B.PublisherID, &B.GenreID, &B.BookTypeID, &B.BooksQty, &B.WritingDate, &B.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return B
}

func SelectAuthorsBook(id int) structures.AuthorBookResponse {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.authorbook WHERE book_id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Abs := structures.AuthorBookResponse{BookID: id}

	for rows.Next() {
		Ab := structures.AuthorBook{}
		err := rows.Scan(&Ab.AuthorID, &Ab.BookID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Abs.AuthorID = append(Abs.AuthorID, Ab.AuthorID)
	}
	return Abs
}

func SelectAuthor(id int) structures.Author {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.author WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	A := structures.Author{}

	for rows.Next() {
		err := rows.Scan(&A.ID, &A.Name, &A.Surname, &A.FatherName, &A.Birthday)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return A
}

func SelectPublisher(id int) structures.Publisher {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.publisher WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	P := structures.Publisher{}

	for rows.Next() {
		err := rows.Scan(&P.ID, &P.Name, &P.Postcode, &P.Address, &P.PhoneNumber, &P.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return P
}

func SelectGenre(id int) structures.Genre {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.genre WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	G := structures.Genre{}

	for rows.Next() {
		err := rows.Scan(&G.ID, &G.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return G
}

func SelectBookType(id int) structures.BookType {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.booktype WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Bt := structures.BookType{}

	for rows.Next() {
		err := rows.Scan(&Bt.ID, &Bt.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return Bt
}

func SelectEvent(id int) structures.Event {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.event WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	E := structures.Event{}

	for rows.Next() {
		err := rows.Scan(&E.ID, &E.Name, &E.RoomID, &E.UserID, &E.EventDate, &E.PeopleQty)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return E
}

func SelectRoom(id int) structures.Room {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.room WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	R := structures.Room{}

	for rows.Next() {
		err := rows.Scan(&R.ID, &R.Name, &R.Capacity)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return R
}

func AddEvent(e *structures.Event) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//Adding new data

	_, err = db.Exec("INSERT INTO public.event (name, room_id, user_id, event_date, people_qty) VALUES ($1, $2, $3, $4, $5)", e.Name, e.RoomID, e.UserID, e.EventDate, e.PeopleQty)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("event was appended successfully")
	}

	defer db.Close()
}

func AddOrder(o *structures.Order) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//Adding new data

	_, err = db.Exec("INSERT INTO public.order (book_exemplar_id, user_id, order_date, refund_date) VALUES ($1, $2, $3, $4)", o.BookExemplarID, o.UserID, o.OrderDate, o.RefundDate)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("order was appended successfully")
	}

	defer db.Close()
}

func DelEv(id int) error {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM public.event WHERE id=%d", id)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("problem with deleting event", err)
		return err
	}
	return nil
}
