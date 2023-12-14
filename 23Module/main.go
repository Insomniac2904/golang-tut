package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello there from module")
	greeter()
	r := mux.NewRouter() // defineing router for	 mux
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r)) // logs if there is some error in serving...and in http.listenandserve we provide port and router
}

func greeter() {
	fmt.Println("FROM MOD")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO GOLANG SERIES</h1>")) //! data from net always go in byte
}
