package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/setalid/snapcar/api/pkg/api"
	"github.com/setalid/snapcar/api/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestIntegration_Rental(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	assert := require.New(t)

	go func() {
		err := api.Run(ctx)
		assert.NoError(err)
	}()

	pickupTime := time.Now()
	pickupRequest := api.RentalPickupRequest{
		BookingNumber:      "1234",
		RegistrationNumber: "AB12345",
		CarCategoryName:    core.SmallCar().Name,
		CustomerSSN:        "12345678901",
		PickupDateTime:     pickupTime,
		PickupMeterReading: 1,
	}

	body, err := json.Marshal(pickupRequest)
	assert.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/rental/pickup", bytes.NewReader(body))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(200, resp.StatusCode)

	returnTime := pickupTime.Add(5 * 24 * time.Hour)
	returnRequest := api.RentalReturnRequest{
		ReturnDateTime:     returnTime,
		ReturnMeterReading: 2, // driven 1km
	}

	body, err = json.Marshal(returnRequest)
	assert.NoError(err)

	req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:8080/rental/return/%s", pickupRequest.BookingNumber), bytes.NewReader(body))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	assert.NoError(err)

	assert.Equal(200, resp.StatusCode)

	var m map[string]any
	err = json.NewDecoder(resp.Body).Decode(&m)
	assert.NoError(err)
	assert.Equal(float64(5), m["price"])
}
