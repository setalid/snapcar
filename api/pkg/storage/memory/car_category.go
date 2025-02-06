package memory

import (
	"context"
	"encoding/json"
	"slices"

	"github.com/setalid/snapcar/api/pkg/core"
)

const (
	CarCategoryKey = "car_category"
)

type CarCategoryRepo struct {
	db *DB
}

func NewCarCategoryRepo(db *DB) *CarCategoryRepo {
	return &CarCategoryRepo{db: db}
}

func (r *CarCategoryRepo) Create(ctx context.Context, c *core.CarCategory) error {
	categories, err := r.getAll(ctx)
	if err != nil {
		return err
	}

	if r.exists(categories, c) {
		return ErrAlreadyExist
	}

	categories = append(categories, *c)

	err = r.setAll(ctx, categories)
	if err != nil {
		return err
	}
	return nil
}

func (r *CarCategoryRepo) Get(ctx context.Context, name string) (*core.CarCategory, error) {
	categories, err := r.getAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, c := range categories {
		if c.Name == name {
			return &c, nil
		}
	}

	return nil, ErrNotFound
}

func (r *CarCategoryRepo) All(ctx context.Context) ([]core.CarCategory, error) {
	all, err := r.getAll(ctx)
	if err != nil {
		return nil, err
	}
	return all, err
}

func (r *CarCategoryRepo) getAll(ctx context.Context) ([]core.CarCategory, error) {
	var out []core.CarCategory

	b, err := r.db.Get(ctx, CarCategoryKey)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(b), &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

func (r *CarCategoryRepo) setAll(ctx context.Context, all []core.CarCategory) error {
	b, err := json.Marshal(all)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, CarCategoryKey, string(b), true)
	if err != nil {
		return err
	}

	return nil
}

func (r *CarCategoryRepo) exists(collections []core.CarCategory, c *core.CarCategory) bool {
	return slices.ContainsFunc(collections, func(x core.CarCategory) bool {
		return c.Name == x.Name
	})
}
