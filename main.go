package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Обработчик маршрута "/"
func home(w http.ResponseWriter, r *http.Request) {
    // Читаем содержимое файла index.html
    html, err := ioutil.ReadFile("index.html")
    if err != nil {
        // Если возникла ошибка, отправляем ошибку 500 и сообщение об ошибке
        http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        return
    }

    // Устанавливаем заголовок Content-Type для ответа как "text/html"
    w.Header().Set("Content-Type", "text/html")

    // Отправляем содержимое страницы index.html в качестве ответа
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
    // Используется функция http.NewServeMux() для инициализации нового рутера, затем
    // функцию "home" регистрируется как обработчик для URL-шаблона "/".
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snip", showSnip)
    mux.HandleFunc("/snip/create", createSnip)

    // Используется функция http.ListenAndServe() для запуска нового веб-сервера. 
    // Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
    // и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
    // мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
    // что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
    log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}