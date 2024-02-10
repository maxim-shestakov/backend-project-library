package structures

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

type UserVer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID             int    `json:"id"`
	BookExemplarID []int  `json:"bookexemplarid"`
	UserID         int    `json:"clientid"`
	OrderDate      string `json:"orderdate"`
	RefundDate     string `json:"refunddate"`
}

type BookExemplar struct {
	ID     int `json:"id"`
	BookID int `json:"bookid"`
	State  int `json:"stateid"`
}

type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PublisherID int    `json:"publisherid"`
	GenreID     int    `json:"genreid"`
	BookTypeID  int    `json:"booktypeid"`
	BooksQty    int    `json:"booksqty"`
	WritingDate string `json:"writingdate"`
	Status      string `json:"status"`
}

type AuthorBook struct {
	AuthorID int `json:"authorid"`
	BookID   int `json:"bookid"`
}

type AuthorBookResponse struct {
	AuthorID []int `json:"authorid"`
	BookID   int `json:"bookid"`
}

type Author struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	FatherName string `json:"fathername"`
	Birthday   string `json:"birthday"`
}

type Publisher struct {
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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	RoomID    int    `json:"roomid"`
	UserID    int    `json:"userid"`
	EventDate string `json:"eventdate"`
	PeopleQty int    `json:"peopleqty"`
}

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type Bucket struct {
	UserID         int   `json:"userid"`
	BookExemplarID []int `json:"bookexemplarid"`
}

//Хранилища:

// Users - хранилище для отправки ответов сервера фронтенду
var Users = map[int]User{}

var Orders = map[int]Order{}

var BookExemplars = map[int]BookExemplar{}

var Books = map[int]Book{}

var AuthorsBook = map[int]AuthorBook{}

var Authors = map[int]Author{}

var Publishers = map[int]Publisher{}

var Genres = map[int]Genre{}

var BookTypes = map[int]BookType{}

var Events = map[int]Event{}

var Rooms = map[int]Room{}

var Buckets = map[int]Bucket{}
