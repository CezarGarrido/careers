package superhero

import (
	"context"

	"github.com/CezarGarrido/careers/entities"
)

type SuperHeroRepo interface {
	Create(ctx context.Context, superHero *entities.SuperHero) (int64, error)
	Update(ctx context.Context, superHero *entities.SuperHero) (*entities.SuperHero, error)
	FindByUUID(ctx context.Context, superUUID int64) (*entities.SuperHero, error)
	Delete(ctx context.Context, superUUID string) error
}
