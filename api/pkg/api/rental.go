package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/setalid/snapcar/api/pkg/core"
	"github.com/setalid/snapcar/api/pkg/utils"
	"go.uber.org/zap"
)

type RentalPickupRequest struct {
	BookingNumber      string    `json:"booking_number"`
	RegistrationNumber string    `json:"registration_number"`
	CarCategoryName    string    `json:"car_category_name"`
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
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		log.Sugar().Infof("Received decoded: %+v", decoded)

		rental := core.NewRental(
			decoded.BookingNumber,
			decoded.RegistrationNumber,
			decoded.CarCategoryName,
			decoded.CustomerSSN,
			decoded.PickupDateTime,
			decoded.PickupMeterReading,
		)

		if err := rental.Validate(r.Context()); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		if err := rentalSvc.RentalPickup(r.Context(), rental); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		encode(w, r, http.StatusOK, map[string]any{})
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

		decoded, err := decode[RentalReturnRequest](r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		update := core.RentalUpdatable{
			ReturnDateTime:     decoded.ReturnDateTime,
			ReturnMeterReading: decoded.ReturnMeterReading,
		}

		if err := update.Validate(r.Context()); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		price, err := rentalSvc.RentalReturn(
			r.Context(),
			bookingNumber,
			update,
		)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
		}

		encode(w, r, http.StatusOK, map[string]any{
			"price": price,
		})
	})
}

func handleGetRentals(
	log *zap.Logger,
	rentalSvc *core.RentalService,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rentals, err := rentalSvc.RentalRepo.All(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
		}

		viewRentals := utils.Map(rentals, func(rental core.Rental) map[string]any {
			return map[string]any{
				"bookingNumber":      rental.BookingNumber,
				"registrationNumber": rental.RegistrationNumber,
				"carCategoryName":    rental.CarCategoryName,
				"pickupDate":         rental.PickupDateTime.String(),
				"pickupMeterReading": fmt.Sprint(rental.PickupMeterReading),
				"returnDate":         rental.ReturnDateTime.String(),
				"returnMeterReading": fmt.Sprint(rental.ReturnMeterReading),
			}
		})

		encode(w, r, http.StatusOK, map[string]any{
			"rentals": viewRentals,
		})
	})
}
