package routes

import (
	"math/big"
	models "models"
	"database/sql"
	http "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	sq "github.com/Masterminds/squirrel"
)

AUTH_ROUTE = "/auth"
USER_ID_KEY = "userId"

func isNeedAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// checking if it is auth route
		if r.URL.Path == AUTH_ROUTE {
			context.Set(r, "needAuth", false)
		} else {
			context.Set(r, "needAuth", true)
		}

		next.ServeHTTP(w, r)
    })
}

func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just setting context
		context.Set(r, USER_ID_KEY, r.Header["Authorization"][0])
		next.ServeHTTP(w, r)
    })
}

func MakeRoutes(db *sql.DB, router *mux.Router) func () () {
	
	router.Use(isNeedAuth)

	router.Use(authorize)

	router.HandleFunc(AUTH_ROUTE, func (w http.ResponseWriter, r *http.Request) {
		// some auth here, for simplicity just store user id in header Authorization
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/photos", func (w http.ResponseWriter, r *http.Request) {
		userIdStr := context.Get(r, USER_ID_KEY)
		userId, err := new(big.Int).SetString(userIdStr, 10)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusOK)
			return
		}

		sql, args := sq.Select(models.Photo.Id.Name, models.Photo.FullPath.Name).
		From(models.Photo.TableName).
		Where(sq.Eq({models.Photo.UserId.Name: userId})).
		ToSql()

		fmt.Println("SQL PHOTOS")
		fmt.Println(sql)
		fmt.Println(args)

		photos := make([]Photo, 0)

		rows, errDb := db.Query(sql, ...args)
		if errDb != nil {
			fmt.Println(errDb)
			w.WriteHeader(http.StatusOK)
			return
		}

		for rows.Next() {
			gotPhoto := models.Photo
			err := rows.Scan(&nowPhoto.Id.Value, &nowPhoto.FullPath.Value)
			if err != nil {
				continue
			}

			append(photos, gotPhoto)
		}

		fmt.Println("ALL")
		fmt.Println(photos)



	}).Methods("GET")

	router.HandleFunc("/photo", func (w http.ResponseWriter, r *http.Request) {
		userIdStr := context.Get(r, USER_ID_KEY)
		userId, err := new(big.Int).SetString(userIdStr, 10)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusOK)
			return
		}

		sq.Insert(models.Photo.TableName).
		Columns(models.Photo.FullPath.Name, models.Photo.PreviewPath.Name, models.Photo.UserId.Name).
		Values()

	}).Methods("POST")


}