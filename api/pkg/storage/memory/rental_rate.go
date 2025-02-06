package memory

import (
	"context"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	GlobalSettingsTable = "global_settings"
	BaseDayRentalKey    = "base_day_rental"
	BaseKmPriceKey      = "base_km_price"
)

type RentalRateRepo struct {
	db *DB
}

func NewRentalRateRepo(db *DB) *RentalRateRepo {
	return &RentalRateRepo{db: db}
}

func (r *RentalRateRepo) Create(ctx context.Context, rate *core.RentalRate) error {
	err := r.db.Set(ctx, GlobalSettingsTable, BaseDayRentalKey, rate.BaseDayRental, false)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, GlobalSettingsTable, BaseKmPriceKey, rate.BaseKmPrice, false)
	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRateRepo) Update(ctx context.Context, update *core.RentalRateUpdatable) error {
	err := r.db.Set(ctx, GlobalSettingsTable, BaseDayRentalKey, update.BaseDayRental, true)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, GlobalSettingsTable, BaseKmPriceKey, update.BaseKmPrice, true)
	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRateRepo) Get(ctx context.Context) (*core.RentalRate, error) {
	var baseDayRental int
	if err := r.db.Get(ctx, GlobalSettingsTable, BaseDayRentalKey, &baseDayRental); err != nil {
		return nil, err
	}

	var baseKmPrice int
	if err := r.db.Get(ctx, GlobalSettingsTable, BaseKmPriceKey, &baseKmPrice); err != nil {
		return nil, err
	}

	return core.NewRentalRate(baseDayRental, baseKmPrice), nil
}
