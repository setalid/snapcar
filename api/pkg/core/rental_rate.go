package core

import "context"

type RentalRateUpdatable struct {
	BaseDayRental int `json:"base_day_rental"`
	BaseKmPrice   int `json:"base_km_price"`
}

type RentalRate struct {
	RentalRateUpdatable
}

func NewRentalRate(baseDayRental, baseKmPrice int) *RentalRate {
	return &RentalRate{
		RentalRateUpdatable: RentalRateUpdatable{
			BaseDayRental: baseDayRental,
			BaseKmPrice:   baseKmPrice,
		},
	}
}

func (r *RentalRate) CalculatePrice(
	baseDayFactor float64,
	baseKmFactor float64,
	numberOfDays int,
	numberOfKm int,
) float64 {
	return float64(r.BaseDayRental)*float64(numberOfDays)*baseDayFactor +
		float64(r.BaseKmPrice)*float64(numberOfKm)*baseKmFactor
}

type RentalRateRepo interface {
	Create(ctx context.Context, r *RentalRate) error
	Update(ctx context.Context, ru *RentalRateUpdatable) error
	Get(ctx context.Context) (*RentalRate, error)
}
