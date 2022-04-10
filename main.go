package main

import (
	"fmt"
	"log"
	"net/http"
)

// r - запрос к серверу
// w - ответ сервера

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {           //
		fmt.Fprintf(w, "ParseForm() err: %v", err)  //
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")                     //Присваиваем переменной name значение(string)
	address := r.FormValue("address")               //Присваиваем переменной name значение(string)
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {                              //Если запрошеный путь не равен ./hello
		http.Error(w, "404 not found", http.StatusNotFound)  //Возвращаем текст "404 not found" со статусом StatusNotFound
		return
	}
	if r.Method != "GET" {                                             //Если метод запроса не равен GET
		http.Error(w, "method is not supported", http.StatusNotFound)  //Возвращаем текст "method is not supported" со статусом StatusNotFound
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)            //Обработка корневой директории(Возвращение статической страници /static/index.html)
	http.HandleFunc("/form", formHandler)   //Обработка перехода по адресу ./form вызывает функцию formHandler
	http.HandleFunc("/hello", helloHandler) //Обработка перехода по адресу ./hello

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {  //Запускаем сервер по адресу localhost:8080, 
		                                                       //получая в переменную err ошибку(В том случае если ошибка возникает, по умолчанию возвращает nil)
															   //В случае если err(Ошибка) не равна nil
		log.Fatal(err)                                         //Прекращаем работу сервера, затем выводим в консоль значение переменной err
	}
}
