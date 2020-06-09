package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Describtion string
	Writer      string
}

type createPostReqest struct {
	Title       string
	Describtion string
	Writer      string
}

type deletePostRequest struct {
	ID          int
	Title       string
	Describtion string
	Writer      string
}

type createPostResponse struct {
	Message string
}

type deletePostResponse struct {
	Message string
}

type readPostResponse struct {
	Posts []Post
}

type readPostRequest struct{}

type fixPostRequest struct {
	ID          int
	Title       string
	Describtion string
	Writer      string
}

type fixPostResponse struct{}

func main() {
	port := 8080
	db := openDB()
	defer db.Close()
	http.Handle("/post", postHandler{db: db})
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

type createHandler struct {
	db *gorm.DB
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request createPostReqest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if request.Title == "" || request.Describtion == "" || request.Writer == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var buf = &Post{
		Title:       request.Title,
		Describtion: request.Describtion,
		Writer:      request.Writer,
	}
	h.db.Create(&buf)
	fmt.Fprint(w, "success")
}

type readHandler struct {
	db *gorm.DB
}

func (h readHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var post []Post
	encoder := json.NewEncoder(w)
	buf := h.db.Find(&post)
	if err := encoder.Encode(&buf); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

type deleteHandler struct {
	db *gorm.DB
}

func (h deleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request deletePostRequest
	var post Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	ID := request.ID
	fmt.Print(h.db.First(&post, ID).Delete(&post))
	h.db.First(&post, ID).Delete(&post)
	fmt.Fprint(w, "success")
}

type fixHandler struct {
	db *gorm.DB
}

func (h fixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request deletePostRequest
	var post Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	ID := request.ID
	buf := h.db.First(&post, ID)
	buf.Update(Post{Writer: request.Writer, Describtion: request.Describtion, Title: request.Title})
}

func openDB() *(gorm.DB) {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to open database")
	}
	db.AutoMigrate(&Post{})
	return db
}

type postHandler struct {
	db *gorm.DB
}

func (h postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		readHandler{db: h.db}.ServeHTTP(w, r)
		break
	case http.MethodPost:
		createHandler{db: h.db}.ServeHTTP(w, r)
		break
	case http.MethodDelete:
		deleteHandler{db: h.db}.ServeHTTP(w, r)
		break
	case http.MethodPut:
		fixHandler{db: h.db}.ServeHTTP(w, r)
		break
	}
}
