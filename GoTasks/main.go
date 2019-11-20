package main

import (
	"GoTasks/Webserver"

	"log"
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
)

func main() {

	//var emri, mbiemri string
	//var mosha int

	//var n, m, nr1, nr2 int
	var s string

	var id, viti int
	var libri, autori, key string
	/*
		fmt.Print("Funksioni faktoril:shenoni numrin : ")
		fmt.Scanln(&n)
		fmt.Println(Factorial.Factorial(n))

		fmt.Print("Funksioni fibonacci:shenoni numrin : ")
		fmt.Scanln(&m)
		fmt.Println(Fibonacci.Fibonacci(m))

		fmt.Print("Funksioni anonim:shenoni parametrat: ")
		fmt.Scan(&nr1)
		fmt.Scan(&nr2)
		value := func() {

			func(a, b int) {
				fmt.Println(nr1 + nr2)
			}(5, 5)
		}
		value()

		fmt.Print("Shenoni te dhenat:emri mbiemri mosha- ")
		fmt.Scan(&emri)
		fmt.Scan(&mbiemri)
		fmt.Scan(&mosha)

		p := &Struct.Person{}

		p.Tedhenatperson(emri, mbiemri, mosha)

		fmt.Println(p.Ndryshomoshen())
	*/
	fmt.Print("Për të thirr web server shenoni Yes ose No : ")
	fmt.Scan(&s)

	fmt.Println("Shenoni Id ")
	fmt.Scan(&id)
	fmt.Println("Shenoni vitin")
	fmt.Scan(&viti)
	fmt.Println("Shenoni librin ")
	fmt.Scan(&libri)
	fmt.Println("Shenoni autorin")
	fmt.Scan(&autori)
	fmt.Println("Shenoni key")
	fmt.Scan(&key)

	b := &Webserver.B

	(*b).Id = id
	(*b).Year = viti
	(*b).Title = libri
	(*b).Author = autori

	fmt.Println(b)

	k := &Webserver.Key

	*k = key

	router := mux.NewRouter()

	if s == "Yes" {

		router.HandleFunc("/", Webserver.Error)
		router.HandleFunc("/book/{string:[0-9]+}", Webserver.Redis)
		router.HandleFunc("/{string:[a-zA-Z0-9]+}", Webserver.Error1)
		router.HandleFunc("/hello/{string:[0-9]+}", Webserver.Error2)
		router.HandleFunc("/hello/{string:[a-zA-Z]+}", Webserver.Hello)

		log.Fatal(http.ListenAndServe(":8080", router))

	}

}
