package memory

import (
	"context"
	"encoding/json"
	"slices"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	RentalKey = "rental"
)

type RentalRepo struct {
	db *DB
}

func NewRentalRepo(db *DB) *RentalRepo {
	return &RentalRepo{db: db}
}

func (r *RentalRepo) Create(ctx context.Context, rental *core.Rental) error {
	rentals, err := r.getAll(ctx)
	if err != nil {
		return err
	}

	if r.exists(rentals, rental) {
		return ErrAlreadyExist
	}

	rentals = append(rentals, *rental)

	err = r.setAll(ctx, rentals)
	if err != nil {
		return err
	}
	return nil
}

func (r *RentalRepo) Update(ctx context.Context, bn string, ru *core.RentalUpdatable) (*core.Rental, error) {
	rentals, err := r.getAll(ctx)
	if err != nil {
		return nil, err
	}

	var out core.Rental
	for i, rental := range rentals {
		if rental.BookingNumber == bn {
			rentals[i].RentalUpdatable = *ru
			out = rentals[i]
		}
	}

	err = r.setAll(ctx, rentals)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *RentalRepo) Get(ctx context.Context, bn string) (*core.Rental, error) {
	rentals, err := r.getAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, c := range rentals {
		if c.BookingNumber == bn {
			return &c, nil
		}
	}

	return nil, ErrNotFound
}

func (r *RentalRepo) All(ctx context.Context) ([]core.Rental, error) {
	all, err := r.getAll(ctx)
	if err != nil {
		return nil, err
	}
	return all, err
}

func (r *RentalRepo) getAll(ctx context.Context) ([]core.Rental, error) {
	var out []core.Rental

	b, err := r.db.Get(ctx, RentalKey)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(b), &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

func (r *RentalRepo) setAll(ctx context.Context, all []core.Rental) error {
	b, err := json.Marshal(all)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, RentalKey, string(b), true)
	if err != nil {
		return err
	}

	return nil
}

func (r *RentalRepo) exists(collections []core.Rental, c *core.Rental) bool {
	return slices.ContainsFunc(collections, func(x core.Rental) bool {
		return c.BookingNumber == x.BookingNumber
	})
}
