package api

import (
	"github.com/gorilla/mux"
	"ic-service/app/api/controller"
	"ic-service/icecream_middleware"
	"log"
	"net/http"
)

const (
	addIcecream    = "AddIcecream"
	deleteIcecream = "DeleteIcecream"
	updateIcecream = "UpdateIcecream"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	v1UnAuthenticatedRouter := r.PathPrefix("/internal/v1").Subrouter()
	v1UnAuthenticatedRouter.HandleFunc("/health_check", controller.GetHeartBeat).Methods(http.MethodGet).Name("GetHeartBeat")

	v1Router := r.PathPrefix("/api/v1").Subrouter()

	//CREATE
	v1Router.HandleFunc("/icecream", controller.AddIcecream).Methods(http.MethodPost).Name(addIcecream)

	//PUT
	v1Router.HandleFunc("/icecream", controller.UpdateIcecream).Methods(http.MethodPut).Name(updateIcecream)

	//DELETE
	v1Router.HandleFunc("/icecream", controller.DeleteIcecream).Methods(http.MethodDelete).Name(deleteIcecream)

	addMiddlewares(v1Router)

	return r
}

func addMiddlewares(routes *mux.Router) {
	log.Print("Adding Middlewares")
	routes.Use(icecream_middleware.GenerateRequestIdHandler)
	routes.Use(icecream_middleware.TokenHandler)
	routes.Use(icecream_middleware.AuthMiddleware)

}
