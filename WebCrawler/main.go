package main

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/0.html", html0)
	router.HandleFunc("/1.html", html1)
	router.HandleFunc("/2.html", html2)
	router.HandleFunc("/3.html", html3)
	router.HandleFunc("/4.html", html4)
	router.HandleFunc("/5.html", html5)
	router.HandleFunc("/6.html", html6)
	router.HandleFunc("/7.html", html7)

	log.Fatal(http.ListenAndServe(":8000", router))

}

func html0(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "0.html")

	c("0.html")

}
func html1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "1.html")
	c("1.html")
}
func html2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "2.html")
	c("2.html")
}
func html3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "3.html")
	c("3.html")
}
func html4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "4.html")
	c("4.html")
}
func html5(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "5.html")
	c("5.html")
}
func html6(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "6.html")
	c("6.html")
}
func html7(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "7.html")
	c("7.html")
}


func c(url string) {
	// Make HTTP request
	response, err := http.Get("http://localhost:8000/" + url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)


}
