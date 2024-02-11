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
	CardID      int     `json:"cardid"`
}

type UserVer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID             int    `json:"id"`
	BookExemplarID int    `json:"bookexemplarid"`
	UserID         int    `json:"clientid"`
	OrderDate      string `json:"orderdate"`
}

type OrderResponse struct {
	ID             int    `json:"id"`
	BookExemplarID []int  `json:"bookexemplarid"`
	UserID         int    `json:"clientid"`
	OrderDate      string `json:"orderdate"`
}

type BookExemplar struct {
	ID      int    `json:"id"`
	BookID  int    `json:"bookid"`
	StateID int    `json:"stateid"`
	Info    string `json:"info"`
}

type Book struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	PublisherID     int    `json:"publisherid"`
	GenreID         int    `json:"genreid"`
	SeriesID        int    `json:"booktypeid"`
	Isbn            string `json:"isbn"`
	PubliishingYear string `json:"publishingyear"`
	Pages           int    `json:"pages"`
	BindingID       int    `json:"binding"`
	Size            string `json:"size"`
	Format          string `json:"format"`
	BookAvailable   bool   `json:"booksqty"`
	Circulation     int    `json:"circulation"`
	Annotation      string `json:"annotation"`
}

type AuthorBook struct {
	AuthorID int `json:"authorid"`
	BookID   int `json:"bookid"`
}

type AuthorBookResponse struct {
	AuthorID []int `json:"authorid"`
	BookID   int   `json:"bookid"`
}

type Author struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	FatherName string `json:"fathername"`
}

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Series struct {
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
	Info      string `json:"info"`
}
type Bindings struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Info     string `json:"info"`
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

var SeriesSt = map[int]Series{}

var Events = map[int]Event{}

var Rooms = map[int]Room{}

var Buckets = map[int]Bucket{}
