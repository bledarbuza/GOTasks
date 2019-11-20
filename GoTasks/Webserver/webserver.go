package Webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

type Books struct {
	Id     int    `json:"id"`
	Year   int    `json:"year"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
	var B Books

	var Key string = ""

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["string"]

	fmt.Fprintln(w, "Rezulati i kthyer nga web server"+res)

}
func Redis(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["string"]


	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(B)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(Key, json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get(res).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "Rezulati i kthyer nga web server"+val)

}

func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Kjo faqe nuk ekziston!")
}

func Error1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["string"]

	if res != "hello" {
		fmt.Fprintln(w, "Pas / duhet te shenohet hello")

	} else if res == "hello" {
		fmt.Fprintln(w, "Pas hello duhet te shenohet nje emer")

	}

}
func Error2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["string"]

	rp, _ := regexp.MatchString("[0-9]+", res)

	if rp == true {
		fmt.Fprintln(w, "Nuk lejohen numra")

	}

}
