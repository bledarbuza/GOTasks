package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Id int          `json:"id"`
	Year int        `json:"year"`
	Author string	`json:"author"`
	Title string	`json:"title"`
}



func main() {

	router := mux.NewRouter()

	router.HandleFunc("/book/{string:[a-zA-Z0-9]+}", Home)

	log.Fatal(http.ListenAndServe(":8080", router))


}

func Home(w http.ResponseWriter , r *http.Request) {

	vars := mux.Vars(r)
	res := vars["string"]

	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Book{Id: 1234, Year: 1999, Author: "Ismail Kadare", Title: "Prilli i Thyer"})
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)

	}
	val, err := client.Get(res).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w,(val))



}