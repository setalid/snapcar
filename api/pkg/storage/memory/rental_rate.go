package memory

import (
	"context"
	"strconv"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	BaseDayRentalKey = "base_day_rental"
	BaseKmPriceKey   = "base_km_price"
)

type RentalRateRepo struct {
	db *DB
}

func NewRentalRateRepo(db *DB) *RentalRateRepo {
	return &RentalRateRepo{db: db}
}

func (r *RentalRateRepo) Create(ctx context.Context, rate *core.RentalRate) error {
	err := r.db.Set(ctx, BaseDayRentalKey, strconv.Itoa(rate.BaseDayRental), false)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, BaseKmPriceKey, strconv.Itoa(rate.BaseKmPrice), false)
	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRateRepo) Update(ctx context.Context, update *core.RentalRateUpdatable) error {
	err := r.db.Set(ctx, BaseDayRentalKey, strconv.Itoa(update.BaseDayRental), true)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, BaseKmPriceKey, strconv.Itoa(update.BaseKmPrice), true)
	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRateRepo) Get(ctx context.Context) (*core.RentalRate, error) {
	baseDayRentalString, err := r.db.Get(ctx, BaseDayRentalKey)
	if err != nil {
		return nil, err
	}

	baseDayRental, err := strconv.Atoi(baseDayRentalString)
	if err != nil {
		return nil, err
	}

	baseKmPriceString, err := r.db.Get(ctx, BaseKmPriceKey)
	if err != nil {
		return nil, err
	}

	baseKmPrice, err := strconv.Atoi(baseKmPriceString)
	if err != nil {
		return nil, err
	}

	return core.NewRentalRate(baseDayRental, baseKmPrice), nil
}
