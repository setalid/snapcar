package core

import (
	"context"
	"fmt"
)

type CarCategory struct {
	Name          string  `json:"name"`
	BaseDayFactor float64 `json:"base_day_factor"`
	BaseKmFactor  float64 `json:"base_km_factor"`
}

func NewCarCategory(name string, baseDayFactor, baseKmFactor float64) *CarCategory {
	return &CarCategory{
		Name:          name,
		BaseDayFactor: baseDayFactor,
		BaseKmFactor:  baseDayFactor,
	}
}

func (c *CarCategory) Validate(_ context.Context) error {
	if c.BaseDayFactor < 0 {
		return fmt.Errorf("expected BaseDayFactor to be >= 0, was: '%f'", c.BaseDayFactor)
	}

	if c.BaseKmFactor < 0 {
		return fmt.Errorf("expected BaseKmFactor to be >= 0, was: '%f'", c.BaseDayFactor)
	}

	return nil
}

func (c *CarCategory) PriceFormula() string {
	var dayFormula string
	if c.BaseDayFactor > 0 {
		dayFormula = fmt.Sprintf("baseDayRental * numberOfDays * %f", c.BaseDayFactor)
	}

	var kmFormula string
	if c.BaseKmFactor > 0 {
		kmFormula = fmt.Sprintf("baseKmPrice * numberOfKm * %f", c.BaseKmFactor)
	}

	if dayFormula != "" && kmFormula != "" {
		return fmt.Sprintf("Price = %s + %s", dayFormula, kmFormula)
	}

	if dayFormula != "" {
		return fmt.Sprintf("Price = %s", dayFormula)
	}

	if kmFormula != "" {
		return fmt.Sprintf("Price = %s", kmFormula)
	}

	return ""
}

type CarCategoryRepo interface {
	Create(ctx context.Context, c *CarCategory) error
	Get(ctx context.Context, name string) (*CarCategory, error)
	All(ctx context.Context) ([]CarCategory, error)
}
