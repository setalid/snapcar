package memory

import (
	"context"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	RentalTable = "rentals"
)

type RentalRepo struct {
	db           *DB
	categoryRepo core.CarCategoryRepo
}

func NewRentalRepo(db *DB, categoryRepo core.CarCategoryRepo) *RentalRepo {
	return &RentalRepo{
		db:           db,
		categoryRepo: categoryRepo,
	}
}

func (r *RentalRepo) Create(ctx context.Context, rental *core.Rental) error {
	_, err := r.categoryRepo.Get(ctx, rental.CarCategoryName)
	if err != nil {
		return err
	}

	return r.db.Set(ctx, RentalTable, rental.BookingNumber, rental, false)
}

func (r *RentalRepo) Update(ctx context.Context, bn string, ru *core.RentalUpdatable) (*core.Rental, error) {
	var rental core.Rental
	err := r.db.Get(ctx, RentalTable, bn, &rental)
	if err != nil {
		return nil, err
	}

	rental.RentalUpdatable = *ru

	return &rental, r.db.Set(ctx, RentalTable, rental.BookingNumber, rental, true)
}

func (r *RentalRepo) Get(ctx context.Context, bn string) (*core.Rental, error) {
	var out core.Rental
	return &out, r.db.Get(ctx, RentalTable, bn, &out)
}

func (r *RentalRepo) All(ctx context.Context) ([]core.Rental, error) {
	var out []core.Rental
	return out, r.db.All(ctx, RentalTable, &out)
}
