package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // разбор аргументов, необходимо вызвать самостоятельно
	fmt.Println(r.Form) // печать данных формы на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v:= range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(r.Form["lol"])
	for l:= range r.Form {/*Ищем нужные нам значения в строке адреса типа
		http://localhost:9090/?url_long=111&url_long=222&lol=517&lol=125 */
		fmt.Println("lol:", l)
		fmt.Println("val:", strings.Join(l, ""))	
	}
	
	fmt.Fprintf(w, "Hello, Oleg!") // отправка данных на клиент
}

func main() {
	http.HandleFunc("/", sayhelloName)       // устанавливаем обработчик
	err := http.ListenAndServe(":9090", nil) // устанавливаем порт, который будем слушать
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
//http://localhost:9090/ Shows "Hello, Oleg!" in browser

/*http://localhost:9090/?url_long=111&url_long=222 send information 
Console:
map[url_long:[111 222] lol:[517 125]]
path /
scheme 
[111 222]
key: url_long
val: 111222
key: lol
val: 517125
map[]
path /favicon.ico
scheme 
[]


*/