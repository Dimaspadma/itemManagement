package main

import (
	"awesomeProject/helper"
	"awesomeProject/sse"
	"awesomeProject/user"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"net/http"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("pgx", "user=postgres password=supersecret host=localhost port=5432 database=postgres sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Init DB once")
	})
}

func main() {

	InitDB()
	err := db.Ping()
	helper.PanicIfError(err)

	// SSE
	sseService := sse.NewService()
	// Goroutine buat manage client connection
	go sseService.ManageClients()
	sseHandler := sse.NewHandler(sseService)

	// User
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService, sseService)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World %s\n", r.URL.Path)
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Hello World %s %s\n", title, page)
	})

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/", Logging(userHandler.GetAllUser)).Methods("GET")
	usersRouter.HandleFunc("/", Logging(userHandler.AddUser)).Methods("POST")
	usersRouter.HandleFunc("/test", Logging(userHandler.TestAddUser)).Methods("POST", "OPTIONS")
	usersRouter.HandleFunc("/{id}", Logging(userHandler.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/events/", sseHandler.BroadcastEvent)

	http.ListenAndServe(":80", r)
}
