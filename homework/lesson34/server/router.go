package server

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson34/server/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {

	h := handlers.NewHandlers(db)
	router.HandleFunc("/user/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/order/", h.GetOrderByID).Methods("GET")
	router.HandleFunc("/product/", h.GetProductByID).Methods("GET")

	router.HandleFunc("/user/", h.CreateUser).Methods("POST")
	router.HandleFunc("POST /product/", h.CreateProduct).Methods("POST")
	router.HandleFunc("POST /order/", h.CreateProduct).Methods("POST")

	router.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELTE")
	router.HandleFunc("/product/", h.DeleteUser).Methods("DELTE")
	router.HandleFunc("/order/", h.DeleteUser).Methods("DELTE")

	router.HandleFunc("/user/", h.UpdateUserById).Methods("PUT")
	router.HandleFunc("/product/", h.DeleteProduct).Methods("PUT")
	router.HandleFunc("/order/", h.DeleteUser).Methods("PUT")

	router.HandleFunc("/users/", h.GetAllUsersByFilter).Methods("GET")
	router.HandleFunc("/products/", h.GetAllProductsByFilter).Methods("GET")

}
