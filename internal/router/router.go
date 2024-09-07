package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knmsh08200/Blog_test/internal/handlers"
	"github.com/knmsh08200/Blog_test/internal/middleware"
)

func NewHandler(ListProvider handlers.BlogMTVYHandlers, IDProvider handlers.BlogIDHandlers) http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/blog/id/{id}", IDProvider.BlogIDHandler).Methods(http.MethodDelete)
	mux.HandleFunc("/blog/list/{id}", ListProvider.BlogListHandler).Methods(http.MethodDelete)

	mux.HandleFunc("/blog/list/", ListProvider.BlogListHandler).Methods(http.MethodGet, http.MethodPost)
	mux.HandleFunc("/blog/id/counter", ListProvider.BlogCounterHandler).Methods(http.MethodGet)
	mux.HandleFunc("/blog/id/", IDProvider.BlogIDHandler).Methods(http.MethodGet, http.MethodPost)
	mux.HandleFunc("/blog/list/find", ListProvider.BlogFindListHandler).Methods(http.MethodGet)

	tokenProtect := middleware.BearerTokenMiddleware(mux, "password")

	handler := middleware.MetricsMiddleware(tokenProtect)
	return handler
}
