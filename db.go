package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_"github.com/go-sql-driver/mysql"
	"net/http"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "crud"
	dbPass = "12345"
	dbHost = "localhost"
	dbPort = "33066"
)

func routers() *chi.Mux {
	router.Get("/posts", Addpost)
	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8",  dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}


type Post struct {
	ID      int    `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}


// CreatePost create a new post
func Addpost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert posts SET title=?, content=?")
	catch(err)

	_, er := query.Exec(post.Title, post.Content)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}


func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}