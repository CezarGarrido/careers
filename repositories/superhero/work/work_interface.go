package work

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type WorkRepo interface {
	Create(ctx context.Context, work *entities.Work) (int64, error)
	FindBySuperID(ctx context.Context, superID int64) (*entities.Work, error)
}
