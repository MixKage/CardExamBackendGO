package routes

import (
	"net/http"

	"github.com/BackendApiCardExam/api/controllers"
	"github.com/gorilla/mux"
)

var RegisterCardRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/server/card", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/v1/server/cardUserAuth", controllers.CreateCardAuthUserOnly).Methods("POST")
	router.HandleFunc("/v1/server/card/{CardId}", controllers.GetCardById).Methods("GET")
	router.HandleFunc("/v1/server/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/v1/server/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/v1/server/ping", controllers.Ping).Methods("GET")
	router.HandleFunc("/v1/server/minimal_version", controllers.RecieveMinimalVerision).Methods("GET")
	router.HandleFunc("/v1/server/questions_help", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/v1/server/universities", controllers.GetUniversities).Methods("GET")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}
