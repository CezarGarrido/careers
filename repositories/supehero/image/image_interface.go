package image

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type ImageRepo interface {
	Create(ctx context.Context, image *entities.Image) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Image, error)
}
