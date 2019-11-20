package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
)
func main(){

	fmt.Println("Starting")

	router := mux.NewRouter()

	router.HandleFunc("/Home/{string:[a-zA-Z]+}",home)
	router.HandleFunc("/Home/{string:[0-9]+}",err1)
	//http.HandleFunc("/{string:[0-9]+}",err1))
	err := http.ListenAndServe(":8081",router)
	if err != nil{
		panic(err)
	}


}
func home(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	res := vars["string"]



	fmt.Println(res)
}
func err1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["string"]

	rp,_ := regexp.MatchString("[0-9]+",res)
	if rp == true{
		fmt.Fprintln(w , "Error")
	}


}