package powerstats

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type PowerstatsRepo interface {
	Create(ctx context.Context, powerstats *entities.Powerstats) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Powerstats, error)
}
