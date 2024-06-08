package server

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/server/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(router *http.ServeMux, db *gorm.DB) {

	h := handlers.NewHandlers(db)
	// router.HandleFunc("GET /user/{id}", h.GetUserByID)
	router.HandleFunc("GET /order/", h.GetOrderByID)
	router.HandleFunc("GET /product/", h.GetProductByID)

	router.HandleFunc("GET /userc/", h.CreateUser)
	router.HandleFunc("POST /product/", h.CreateProduct)
	router.HandleFunc("POST /order/", h.CreateOrder)

	router.HandleFunc("DELETE /user/{id}", h.DeleteUser)
	router.HandleFunc("DELETE /product/", h.DeleteUser)
	router.HandleFunc("DELETE /order/", h.DeleteUser)

	// router.HandleFunc("POST /user/", h.UpdateUserById)
	router.HandleFunc("PUT /product/", h.DeleteProduct)
	router.HandleFunc("PUT /order/", h.DeleteUser)

	router.HandleFunc("GET /users/", h.GetAllUsersByFilter)
}
