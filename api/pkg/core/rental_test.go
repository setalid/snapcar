package core

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRentalUpdatable_Validate(t *testing.T) {
	tests := []struct {
		name          string
		rental        *RentalUpdatable
		expectedError bool
	}{
		{
			name:          "Positive return meter reading",
			rental:        &RentalUpdatable{ReturnMeterReading: 100},
			expectedError: false,
		},
		{
			name:          "Negative return meter reading",
			rental:        &RentalUpdatable{ReturnMeterReading: -1},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			err := tt.rental.Validate(context.Background())
			if tt.expectedError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestRental_Validate(t *testing.T) {
	tests := []struct {
		name          string
		rental        *Rental
		expectedError bool
	}{
		{
			name: "Valid Rental",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "AB12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550678",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: false,
		},
		{
			name: "Missing BookingNumber",
			rental: &Rental{
				BookingNumber:      "",
				RegistrationNumber: "AB12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550678",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: true,
		},
		{
			name: "Missing registration number",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550678",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: true,
		},
		{
			name: "Invalid registration number",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "A12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550678",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: true,
		},
		{
			name: "Missing customer ssn",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "AB12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: true,
		},
		{
			name: "Invalid customer ssn",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "AB12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550",
				PickupDateTime:     time.Now(),
				PickupMeterReading: 5000,
			},
			expectedError: true,
		},
		{
			name: "Meter reading less than 0",
			rental: &Rental{
				BookingNumber:      "booking-1",
				RegistrationNumber: "AB12345",
				CarCategoryName:    "Small car",
				CustomerSSN:        "01019550678",
				PickupDateTime:     time.Now(),
				PickupMeterReading: -1,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			err := tt.rental.Validate(context.Background())
			if tt.expectedError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
