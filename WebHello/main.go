package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
func main(){

	fmt.Println("Starting")
	http.HandleFunc("/Home/{string:[a-zA-z]+}",home)
	err := http.ListenAndServe(":8081",nil)
	if err != nil{
		panic(err)
	}


}
func home(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	res := vars["string"]

	fmt.Println(res)
}