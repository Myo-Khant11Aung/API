package main

import "net/http"

func main(){
	db := NewDB("postgres", "postgres://postgres:221104@127.0.0.1:5432/tasks?sslmode=disable")
	router := NewRouter(db)

	server := http.Server{
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}