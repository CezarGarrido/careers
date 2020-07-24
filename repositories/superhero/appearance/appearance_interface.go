package appearance

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type AppearanceRepo interface {
	Create(ctx context.Context, appearance *entities.Appearance) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Appearance, error)
}
