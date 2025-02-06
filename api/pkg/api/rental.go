package api

import (
	"net/http"
	"time"

	"github.com/setalid/snapcar/api/pkg/core"
	"go.uber.org/zap"
)

type RentalPickupRequest struct {
	BookingNumber      string    `json:"booking_number"`
	CarCategoryName    string    `json:"car_category_id"`
	CustomerSSN        string    `json:"customer_ssn"`
	PickupDateTime     time.Time `json:"pickup_date_time"`
	PickupMeterReading int       `json:"pickup_meter_reading"`
}

func handleRentalPickup(
	log *zap.Logger,
	rentalSvc *core.RentalService,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoded, err := decode[RentalPickupRequest](r)
		if err != nil {
			log.Error(err.Error())
			return
		}

		rental := core.NewRental(
			decoded.BookingNumber,
			decoded.CarCategoryName,
			decoded.CustomerSSN,
			decoded.PickupDateTime,
			decoded.PickupMeterReading,
		)

		if err := rental.Validate(r.Context()); err != nil {
			log.Error(err.Error())
			return
		}

		if err := rentalSvc.RentalRepo.Create(r.Context(), rental); err != nil {
			log.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

type RentalReturnRequest struct {
	ReturnDateTime     time.Time `json:"return_date_time"`
	ReturnMeterReading int       `json:"return_meter_reading"`
}

func handleRentalReturn(
	log *zap.Logger,
	rentalSvc *core.RentalService,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookingNumber := r.PathValue("bn")
		if bookingNumber == "" {
			log.Error("missing booking number in request path")
		}

		decoded, err := decode[RentalReturnRequest](r)
		if err != nil {
			log.Error(err.Error())
			return
		}

		update := core.RentalUpdatable{
			ReturnDateTime:     decoded.ReturnDateTime,
			ReturnMeterReading: decoded.ReturnMeterReading,
		}

		if err := update.Validate(r.Context()); err != nil {
			log.Error(err.Error())
			return
		}

		price, err := rentalSvc.RentalReturn(
			r.Context(),
			bookingNumber,
			update,
		)
		if err != nil {
			log.Error(err.Error())
		}

		encode(w, r, http.StatusOK, map[string]any{
			"price": price,
		})
	})
}
