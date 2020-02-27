package routes

import (
	models "models"
	"database/sql"
	http "net/http"
	"github.com/gorilla/mux"
)

func MakeRoutes(*sql.DB, *http) func () () {
	http.HandleFunc()
}