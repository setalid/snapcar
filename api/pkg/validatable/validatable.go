package validatable

import "context"

type Validatable interface {
	Validate(ctx context.Context) error
}
