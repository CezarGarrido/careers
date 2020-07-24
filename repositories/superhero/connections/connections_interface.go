package connections

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type ConnectionsRepo interface {
	Create(ctx context.Context, connections *entities.Connections) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Connections, error)
}
