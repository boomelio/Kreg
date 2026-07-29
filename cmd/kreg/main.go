package main

import (
	"fmt"
	"github.com/Dreamstick9/Kreg/internal/config"
	"github.com/Dreamstick9/Kreg/internal/repository"
	"log"
	"net/http"
)

func homepagehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world it's me kreg! I'm alive yay"))
	fmt.Println(r.Method, r.URL.Path)
}
func secretpagehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ayo why you here boi, go back"))
	fmt.Println(r.Method, r.URL.Path)
}
func main() {
	var cfg = config.Load()
	var db, err = repository.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("PRAGMA journal_mode = WAl")
	if err != nil {
		log.Fatal(err)
	}

	err = repository.CreateSchema(db)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homepagehandler)
	http.HandleFunc("/secret", secretpagehandler)
	http.ListenAndServe(cfg.ServerPort, nil)
}
