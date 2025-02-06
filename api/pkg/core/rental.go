package core

import (
	"context"
	"math"
	"time"
)

type RentalUpdatable struct {
	ReturnDateTime     time.Time `json:"return_date_time"`
	ReturnMeterReading int       `json:"return_meter_reading"`
}

func (c *RentalUpdatable) Validate(_ context.Context) error {
	return nil
}

type Rental struct {
	BookingNumber      string    `json:"booking_number"`
	CarCategoryName    string    `json:"car_category_name"`
	CustomerSSN        string    `json:"customer_ssn"`
	PickupDateTime     time.Time `json:"pickup_date_time"`
	PickupMeterReading int       `json:"pickup_meter_reading"`
	RentalUpdatable
}

func NewRental(
	bookingNumber string,
	carCategoryID string,
	customerSSN string,
	pickupDateTime time.Time,
	pickupMeterReading int,
) *Rental {
	return &Rental{
		BookingNumber:      bookingNumber,
		CarCategoryName:    carCategoryID,
		CustomerSSN:        customerSSN,
		PickupDateTime:     pickupDateTime,
		PickupMeterReading: pickupMeterReading,
	}
}

func (c *Rental) Validate(_ context.Context) error {
	return nil
}

type RentalRepo interface {
	Create(ctx context.Context, r *Rental) error
	Update(ctx context.Context, bn string, ru *RentalUpdatable) (*Rental, error)
	Get(ctx context.Context, bn string) (*Rental, error)
	All(ctx context.Context) ([]Rental, error)
}

type RentalService struct {
	RentalRepo      RentalRepo
	CarCategoryRepo CarCategoryRepo
	RentalRateRepo  RentalRateRepo
}

func NewRentalService(
	rentalRepo RentalRepo,
	carCategoryRepo CarCategoryRepo,
	rentalRateRepo RentalRateRepo,
) *RentalService {
	return &RentalService{
		RentalRepo:      rentalRepo,
		CarCategoryRepo: carCategoryRepo,
		RentalRateRepo:  rentalRateRepo,
	}
}

func (s *RentalService) RentalReturn(ctx context.Context, bn string, update RentalUpdatable) (float64, error) {
	rental, err := s.RentalRepo.Update(ctx, bn, &update)
	if err != nil {
		return -1, err
	}

	category, err := s.CarCategoryRepo.Get(ctx, rental.CarCategoryName)
	if err != nil {
		return -1, err
	}

	rate, err := s.RentalRateRepo.Get(ctx)
	if err != nil {
		return -1, err
	}

	timediff := update.ReturnDateTime.Sub(rental.PickupDateTime)
	days := int(math.Ceil(timediff.Hours() / 24)) // always round up to nearest day

	price := rate.CalculatePrice(category.BaseDayFactor, category.BaseKmFactor, days, update.ReturnMeterReading)
	return price, nil
}
