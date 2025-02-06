package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRentalRate_CalculatePrice(t *testing.T) {
	tests := []struct {
		name         string
		rate         *RentalRate
		category     *CarCategory
		numberOfDays int
		numberOfKm   int
		expects      float64
	}{
		{
			name:         "1 * 5 * 1 + 1 * 5 * 0",
			rate:         NewRentalRate(1, 1),
			category:     SmallCar(),
			numberOfDays: 5,
			numberOfKm:   5,
			expects:      float64(5),
		},
		{
			name:         "1 * 5 * 1.3 + 1 * 5 * 1",
			rate:         NewRentalRate(1, 1),
			category:     Combi(),
			numberOfDays: 5,
			numberOfKm:   5,
			expects:      float64(11.5),
		},
		{
			name:         "1 * 5 * 1.5 + 1 * 5 * 1.5",
			rate:         NewRentalRate(1, 1),
			category:     Truck(),
			numberOfDays: 5,
			numberOfKm:   5,
			expects:      float64(15),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			price := tt.rate.CalculatePrice(
				tt.category.BaseDayFactor,
				tt.category.BaseKmFactor,
				tt.numberOfDays,
				tt.numberOfKm,
			)
			assert.Equal(tt.expects, price)
		})
	}
}
