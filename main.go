package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Обработчик маршрута "/"
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		html, err := ioutil.ReadFile("404.html")
        if err != nil {
            http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        }
       w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Write(html)
        return
    } 
        
    html, err := ioutil.ReadFile("index.html")
    if err != nil {
        http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    
    w.Write(html)
}


// Обработчик маршрута /snip/
func showSnip(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("№1  Первая заметка \n№2  Вторая заметка\n№3  Помыть посуду"))
}

// Обработчик маршрута /snip/create/
func createSnip(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Сервис для создания заметок\nПока у вас нет ни одной заметки"))
}

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snip", showSnip)
    mux.HandleFunc("/snip/create", createSnip)

    log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}