package main

import (
    "log"
    "net/http"
    "fmt"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func main() {
    fmt.Println("Start")
    router := mux.NewRouter()
    router.HandleFunc("/signin",abc).Methods("POST")

    fmt.Println("Listen and Server")
    hanxd:=handlers.AllowedOrigins([]string{"*"})
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(hanxd)(router)))
}

func abc(rw http.ResponseWriter,rq *http.Request){
    rw.Write([]byte("welcome to kloudone!"))
}