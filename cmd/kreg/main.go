package main

import (
	"fmt"
	"net/http"
	"github.com/Dreamstick9/Kreg/internal/config"
)

var cfg = config.Load()


func homepagehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world it's me kreg! I'm alive yay"))
	fmt.Println(r.Method, r.URL.Path)
}
func secretpagehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ayo why you here boi, go back"))
	fmt.Println(r.Method, r.URL.Path)
}

func main() {
	http.HandleFunc("/", homepagehandler)
	http.HandleFunc("/secret", secretpagehandler)
	http.ListenAndServe(cfg.ServerPort, nil)
}
