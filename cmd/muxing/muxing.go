package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{id:[\\w|\\W]+}", helloHandler).Methods("GET")
	router.HandleFunc("/bad", badHandler).Methods("GET")
	router.HandleFunc("/data", dataHandler).Methods("POST")
	router.HandleFunc("/headers", headersHandler).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func helloHandler(resp http.ResponseWriter, r *http.Request) {
	_, err := resp.Write([]byte("Hello," + strings.TrimPrefix(r.URL.Path, "/name/")))
	if err != nil {
		return
	}
	resp.WriteHeader(http.StatusOK)
}

func badHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(resp http.ResponseWriter, r *http.Request) {
	bodyB, _ := ioutil.ReadAll(r.Body)
	_, err := resp.Write([]byte("I got message:\n" + string(bodyB)))
	if err != nil {
		return
	}
	resp.WriteHeader(http.StatusOK)
}

func headersHandler(resp http.ResponseWriter, r *http.Request) {
	aValue, _ := strconv.Atoi(r.Header.Get("a"))
	bValue, _ := strconv.Atoi(r.Header.Get("b"))
	resp.Header().Add("A+B", strconv.Itoa(aValue+bValue))
	resp.WriteHeader(http.StatusOK)
}
