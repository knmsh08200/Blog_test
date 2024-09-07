package handlers

import (
	"net/http"
)

type BlogMTVYHandlers interface { // где это .... используется
	BlogListHandler(w http.ResponseWriter, r *http.Request)
	BlogCounterHandler(w http.ResponseWriter, r *http.Request)
	BlogFindListHandler(w http.ResponseWriter, r *http.Request)
}

type BlogIDHandlers interface {
	BlogIDHandler(w http.ResponseWriter, r *http.Request)
}

// перенсти сюда функции  структуры BlogHandler
