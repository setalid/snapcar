package api

import (
	"net/http"

	"github.com/setalid/snapcar/api/pkg/core"
	"go.uber.org/zap"
)

func NewRootHandler(
	log *zap.Logger,
	rentalSvc *core.RentalService,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /rental/pickup", handleRentalPickup(log, rentalSvc))
	mux.Handle("POST /rental/return/{bn}", handleRentalReturn(log, rentalSvc))

	var handler http.Handler = mux
	return handler
}
