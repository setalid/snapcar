package api

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/setalid/snapcar/api/pkg/core"
	"go.uber.org/zap"
)

func NewRootHandler(
	log *zap.Logger,
	carCategoryRepo core.CarCategoryRepo,
	rentalSvc *core.RentalService,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /rental/pickup", handleRentalPickup(log, rentalSvc))
	mux.Handle("POST /rental/return/{bn}", handleRentalReturn(log, rentalSvc))
	mux.Handle("GET /rental/all", handleGetRentals(log, rentalSvc))
	mux.Handle("GET /category/all", handleGetCarCategories(log, carCategoryRepo))

	var handler http.Handler = mux

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	handler = c.Handler(handler)

	return handler
}
