package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
    
    if r.URL.Path != "/" {
		showNotFoundError(w, r)
		return
	}
    
 	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
    
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showNotFoundError(w http.ResponseWriter, r *http.Request) {
    ts, err := template.ParseFiles("./ui/html/error404.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик маршрута /snip/
func showSnip(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        showNotFoundError(w, r)
        return
    }
    fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

// Обработчик для создания новой заметки.
func createSnip(w http.ResponseWriter, r *http.Request) {
    // Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
    // что http.MethodPost является строкой и содержит текст "POST".
    if r.Method != http.MethodPost {
        // Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
        // и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
        // Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
        w.WriteHeader(405)
        w.Write([]byte("GET-Метод запрещен!"))
        return
    }
 
    w.Write([]byte("Создание новой заметки..."))
}