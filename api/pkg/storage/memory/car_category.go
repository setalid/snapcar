package memory

import (
	"context"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	CarCategoryTable = "car_categories"
)

type CarCategoryRepo struct {
	db *DB
}

func NewCarCategoryRepo(db *DB) *CarCategoryRepo {
	return &CarCategoryRepo{db: db}
}

func (r *CarCategoryRepo) Create(ctx context.Context, c *core.CarCategory) error {
	return r.db.Set(ctx, CarCategoryTable, c.Name, c, false)
}

func (r *CarCategoryRepo) Get(ctx context.Context, name string) (*core.CarCategory, error) {
	var out core.CarCategory
	return &out, r.db.Get(ctx, CarCategoryTable, name, &out)
}

func (r *CarCategoryRepo) All(ctx context.Context) ([]core.CarCategory, error) {
	var out []core.CarCategory
	return out, r.db.All(ctx, CarCategoryTable, &out)
}
