package server

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/server/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(router *http.ServeMux, db *gorm.DB) {
	h := handlers.NewHandlers(db)
	router.HandleFunc("GET /user/", h.GetUserByID)
	router.HandleFunc("/order/", h.GetOrderByID)
	router.HandleFunc("/product/", h.GetProductByID)
}
