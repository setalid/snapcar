package core

import (
	"context"
	"errors"
	"math"
	"regexp"
	"time"
)

var (
	registrationNumberRegex *regexp.Regexp
	ssnRegex                *regexp.Regexp
)

func init() {
	var err error
	registrationNumberRegex, err = regexp.Compile(`^[A-Za-z]{2}\d{5}$`)
	if err != nil {
		panic(err)
	}

	ssnRegex, err = regexp.Compile(`^\d{11}$`)
	if err != nil {
		panic(err)
	}
}

type RentalUpdatable struct {
	ReturnDateTime     time.Time `json:"return_date_time"`
	ReturnMeterReading int       `json:"return_meter_reading"`
}

func (r *RentalUpdatable) Validate(_ context.Context) error {
	if r.ReturnMeterReading < 0 {
		return errors.New("return meter reading must be a positive number")
	}
	return nil
}

type Rental struct {
	BookingNumber      string    `json:"booking_number"`
	RegistrationNumber string    `json:"registration_number"`
	CustomerSSN        string    `json:"customer_ssn"`
	CarCategoryName    string    `json:"car_category_name"`
	PickupDateTime     time.Time `json:"pickup_date_time"`
	PickupMeterReading int       `json:"pickup_meter_reading"`
	RentalUpdatable
}

func NewRental(
	bookingNumber string,
	registrationNumber string,
	carCategoryID string,
	customerSSN string,
	pickupDateTime time.Time,
	pickupMeterReading int,
) *Rental {
	return &Rental{
		BookingNumber:      bookingNumber,
		RegistrationNumber: registrationNumber,
		CarCategoryName:    carCategoryID,
		CustomerSSN:        customerSSN,
		PickupDateTime:     pickupDateTime,
		PickupMeterReading: pickupMeterReading,
	}
}

func (r *Rental) Validate(_ context.Context) error {
	if r.BookingNumber == "" {
		return errors.New("booking number is required")
	}

	if r.RegistrationNumber == "" {
		return errors.New("regiration number is required")
	}

	if valid := registrationNumberRegex.MatchString(r.RegistrationNumber); !valid {
		return errors.New("registration number must be in the form 'AB12345'")
	}

	if r.CustomerSSN == "" {
		return errors.New("customer SSN is required")
	}

	if valid := ssnRegex.MatchString(r.CustomerSSN); !valid {
		return errors.New("customer ssn number must be in the form 'DDMMYYXXXXX'")
	}

	if r.PickupMeterReading < 0 {
		return errors.New("meter reading must be a positive number")
	}
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
