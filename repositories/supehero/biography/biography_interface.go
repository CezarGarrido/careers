package biography

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type BiographyRepo interface {
	Create(ctx context.Context, biography *entities.Biography) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Biography, error)
}
