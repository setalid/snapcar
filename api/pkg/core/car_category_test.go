package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCarCategory_Validate(t *testing.T) {
	tests := []struct {
		name          string
		category      *CarCategory
		expectedError bool
	}{
		{
			name:          "Valid CarCategory",
			category:      NewCarCategory("Valid", 1, 1),
			expectedError: false,
		},
		{
			name:          "Negative BaseDayFactor",
			category:      NewCarCategory("Invalid Day Factor", -1, 1),
			expectedError: true,
		},
		{
			name:          "Negative BaseKmFactor",
			category:      NewCarCategory("Invalid Km Factor", 1, -1),
			expectedError: true,
		},
		{
			name:          "Negative Both Factors",
			category:      NewCarCategory("Invalid Both", -1, -1),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			err := tt.category.Validate(context.Background())
			if tt.expectedError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestCarCategory_PriceFormula(t *testing.T) {
	tests := []struct {
		category *CarCategory
		expects  string
	}{
		{
			category: SmallCar(),
			expects:  "Price = baseDayRental * numberOfDays * 1.0",
		},
		{
			category: Combi(),
			expects:  "Price = baseDayRental * numberOfDays * 1.3 + baseKmPrice * numberOfKm * 1.0",
		},
		{
			category: Truck(),
			expects:  "Price = baseDayRental * numberOfDays * 1.5 + baseKmPrice * numberOfKm * 1.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.category.Name, func(t *testing.T) {
			assert := require.New(t)
			formula := tt.category.PriceFormula()
			assert.Equal(tt.expects, formula)
		})
	}
}
