package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"encoding/json"

	"library_project/server/postgresql"
	st "library_project/server/structures"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret = []byte("gBElG5NThZSye")
)

func ReturnToken(w http.ResponseWriter, r *http.Request) {
	var user st.UserVer
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := jwt.MapClaims{
		"login":          user.Email,
		"hashedpassword": user.Password,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	SignedToken, _ := jwtToken.SignedString(secret)
	tk := fmt.Sprintf("Bearer %s", SignedToken)
	w.Header().Set("Authorization", tk)
	w.WriteHeader(http.StatusOK)
}

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	st.Users = postgresql.SelectAllUsers()
// 	resp, err := json.Marshal(st.Users)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user st.User
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postgresql.AddUser(&user)
	w.WriteHeader(http.StatusCreated)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var ordresp st.OrderResponse
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if id == user_id {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ordresp = postgresql.SelectAllOrders(id)
		resp, err := json.Marshal(ordresp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

//	func VerifyUser(w http.ResponseWriter, r *http.Request) {
//		var user st.UserVer
//		user.Password = r.URL.Query().Get("password")
//		user.Email = r.URL.Query().Get("email")
//		UserResp := postgresql.SelectUserData(&user)
//		resp, err := json.Marshal(UserResp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		w.Write(resp)
//	}

func verifyUser(token string) (bool, int) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Printf("Failed to parse token: %s\n", err)
		return false, -1
	}
	if !jwtToken.Valid {
		return false, -1
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, -1
	}
	loginRaw, ok := claims["login"]
	if !ok {
		return false, -1
	}
	login, ok := loginRaw.(string)
	if !ok {
		return false, -1
	}
	passwordRaw, ok := claims["hashedpassword"]
	if !ok {
		return false, -1
	}
	password, ok := passwordRaw.(string)
	if !ok {
		return false, -1
	}
	id, err := postgresql.SelectVerUser(login, password)
	if err != nil {
		return false, -1
	}
	return true, id
}

func GetBookEx(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	BookExResp := postgresql.SelectBookEx(id)
	resp, err := json.Marshal(BookExResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	BookResp := postgresql.SelectBook(id)
	resp, err := json.Marshal(BookResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	BooksResp := postgresql.SelectBooks()
	resp, err := json.Marshal(BooksResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetAuthorBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	AuthorsBookResp := postgresql.SelectAuthorsBook(id)
	resp, err := json.Marshal(AuthorsBookResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	AuthorResp := postgresql.SelectAuthor(id)
	resp, err := json.Marshal(AuthorResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetPublisher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	PublisherResp := postgresql.SelectPublisher(id)
	resp, err := json.Marshal(PublisherResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	GenreResp := postgresql.SelectGenre(id)
	resp, err := json.Marshal(GenreResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetSeries(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	BtResp := postgresql.SelectSeries(id)
	resp, err := json.Marshal(BtResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if id == user_id {
		EvResp := postgresql.SelectEvent(id)
		resp, err := json.Marshal(EvResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	RoomResp := postgresql.SelectRoom(id)
	resp, err := json.Marshal(RoomResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func PostEvent(w http.ResponseWriter, r *http.Request) {
	var event st.Event
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if event.UserID == user_id {
		postgresql.AddEvent(&event)
		w.WriteHeader(http.StatusCreated)
	}
}

func PostOrder(w http.ResponseWriter, r *http.Request) {
	var order st.Order
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if order.UserID == user_id {
		postgresql.AddOrder(&order)
		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if id == user_id {
		err = postgresql.DelEv(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PostBasket(w http.ResponseWriter, r *http.Request) {
	var Basket st.Basket
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &Basket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if Basket.UserID == user_id {
		st.Baskets[Basket.UserID] = Basket
		w.WriteHeader(http.StatusCreated)
	}
}

func GetBasket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	authHeader := r.Header.Get("Authorization")        //получение токена с типом данных
	token := strings.TrimPrefix(authHeader, "Bearer ") //отделение токена от типа данных
	_, user_id := verifyUser(token)                    //соотношение токена с БД
	if id == user_id {
		resp, err := json.Marshal(st.Baskets[id])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func GetBinding(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	BResp := postgresql.SelectBindings(id)
	resp, err := json.Marshal(BResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
