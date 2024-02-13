package postgresql

import (
	"database/sql"
	"fmt"

	"library_project/server/structures"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db, err
}

// SelectAllUsers возвращает мапу из базы данных, которая заполняется записями о пользователях из неё
func SelectAllUsers() map[int]structures.User {
	db, _ := Connection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM postgres.public.users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := map[int]structures.User{}

	for rows.Next() {
		u := structures.User{}

		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.FatherName, &u.PhoneNumber, &u.Email, &u.Birthday, &u.CardID, &u.Rating, &u.Password)
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
	db, _ := Connection()
	//Adding new data

	_, err := db.Exec("INSERT INTO public.users (name, surname, father_name, phone_number, email, birthday, card_id, rating, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", u.Name, u.Surname, u.FatherName, u.PhoneNumber, u.Email, u.Birthday, u.CardID, u.Rating, u.Password)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("user was appended successfully")
	}

	defer db.Close()
}

func SelectUserData(u *structures.UserVer) structures.User {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.users WHERE password='%s' AND email='%s'", u.Password, u.Email)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	us := structures.User{}
	for rows.Next() {
		err := rows.Scan(&us.ID, &us.Name, &us.Surname, &us.FatherName, &us.PhoneNumber, &us.Email, &us.Birthday, &us.CardID, &us.Rating, &us.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return us
}

func SelectAllOrders(id int) structures.OrderResponse {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.orders WHERE user_id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Orders := structures.OrderResponse{}

	for rows.Next() {
		o := structures.Order{}
		err := rows.Scan(&o.ID, &o.BookExemplarID, &o.UserID, &o.OrderDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Orders.BookExemplarID = append(Orders.BookExemplarID, o.BookExemplarID)
	}
	return Orders
}

func SelectBookEx(id int) structures.BookExemplar {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.bookexemplar WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	Bx := structures.BookExemplar{}
	for rows.Next() {
		err := rows.Scan(&Bx.ID, &Bx.BookID, &Bx.StateID, &Bx.Info)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return Bx
}

func SelectBook(id int) structures.Book {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.books WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	B := structures.Book{}

	for rows.Next() {
		err := rows.Scan(&B.ID, &B.Name, &B.PublisherID, &B.GenreID, &B.SeriesID, &B.Isbn, &B.PubliishingYear, &B.Pages, &B.BindingID, &B.Size, &B.Format, &B.Circulation, &B.Annotation)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return B
}

func SelectBooks() map[int]structures.Book {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.books")
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Bm := map[int]structures.Book{}

	for rows.Next() {
		B := structures.Book{}
		err := rows.Scan(&B.ID, &B.Name, &B.PublisherID, &B.GenreID, &B.SeriesID, &B.Isbn, &B.PubliishingYear, &B.Pages, &B.BindingID, &B.Size, &B.Format, &B.Circulation, &B.Annotation)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Bm[B.ID] = B
	}
	return Bm
}

func SelectAuthorsBook(id int) structures.AuthorBookResponse {
	db, _ := Connection()
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
		err := rows.Scan(&Ab.BookID, &Ab.AuthorID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Abs.AuthorID = append(Abs.AuthorID, Ab.AuthorID)
	}
	return Abs
}

func SelectAuthor(id int) structures.Author {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.authors WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	A := structures.Author{}

	for rows.Next() {
		err := rows.Scan(&A.ID, &A.Name, &A.Surname, &A.FatherName)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return A
}

func SelectPublisher(id int) structures.Publisher {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.publishers WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	P := structures.Publisher{}

	for rows.Next() {
		err := rows.Scan(&P.ID, &P.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return P
}

func SelectGenre(id int) structures.Genre {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.genres WHERE id=%d", id)
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

func SelectSeries(id int) structures.Series {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.series WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	S := structures.Series{}

	for rows.Next() {
		err := rows.Scan(&S.ID, &S.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return S
}

func SelectEvent(id int) structures.Event {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.events WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	E := structures.Event{}

	for rows.Next() {
		err := rows.Scan(&E.ID, &E.Name, &E.RoomID, &E.UserID, &E.EventDate, &E.PeopleQty, &E.Info)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return E
}

func SelectRoom(id int) structures.Room {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.rooms WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	R := structures.Room{}

	for rows.Next() {
		err := rows.Scan(&R.ID, &R.Name, &R.Capacity, &R.Info)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return R
}

func AddEvent(e *structures.Event) {
	db, _ := Connection()
	//Adding new data

	_, err := db.Exec("INSERT INTO public.events (name, room_id, user_id, event_date, people_count, info) VALUES ($1, $2, $3, $4, $5, $6)", e.Name, e.RoomID, e.UserID, e.EventDate, e.PeopleQty, e.Info)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("event was appended successfully")
	}

	defer db.Close()
}

func AddOrder(o *structures.Order) {
	db, _ := Connection()
	//Adding new data

	_, err := db.Exec("INSERT INTO public.orders (book_exemplar_id, user_id, order_date) VALUES ($1, $2, $3)", o.BookExemplarID, o.UserID, o.OrderDate)
	if err != nil {
		fmt.Printf("addUser: %v\n", err)
	} else {
		fmt.Println("order was appended successfully")
	}

	defer db.Close()
}

func DelEv(id int) error {
	db, _ := Connection()
	query := fmt.Sprintf("DELETE FROM public.events WHERE id=%d", id)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("problem with deleting event", err)
		return err
	}
	return nil
}

func SelectBindings(id int) structures.Bindings {
	db, _ := Connection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM postgres.public.bindings WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	B := structures.Bindings{}

	for rows.Next() {
		err := rows.Scan(&B.ID, &B.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return B
}
