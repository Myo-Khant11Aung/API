package main

import (
	"database/sql"
	"net/http"
)

type Router struct{
    routes map[string]map[string]http.HandlerFunc
}

func (r *Router) addRoute(method string, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request){
	EnableCors(w, req)

	if handlers, ok := r.routes[req.URL.Path]; ok{
		if handler, methodExists := handlers[req.Method]; methodExists{
			handler(w,req)
			return
		}
	}
	http.NotFound(w, req)
}

func NewRouter(db *sql.DB) *Router{
	r := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
    service := NewService(db)
	r.addRoute("POST", "/tasks", service.AddTask)
	r.addRoute("GET", "/tasks", service.GetTasks)
	r.addRoute("DELETE", "/tasks", service.DeletTask)
	return r
}