package superhero

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CezarGarrido/careers/entities"
)

var ErrSuperHeroNotFound = errors.New("Super hero not found")

func NewPgSQLSuperHeroRepo(Conn *sql.DB) SuperHeroRepo {
	return &postgresSuperHeroRepo{Conn: Conn}
}

type postgresSuperHeroRepo struct {
	Conn *sql.DB
}

func (this *postgresSuperHeroRepo) Create(ctx context.Context, superHero *entities.SuperHero) (int64, error) {

	query := `INSERT INTO super_heroes (uuid, super_hero_api_id, name, created_at) VALUES($1,$2,$3,$4) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		superHero.UUID,
		superHero.SuperHeroApiID,
		superHero.Name,
		superHero.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresSuperHeroRepo) Update(ctx context.Context, superHero *entities.SuperHero) (*entities.SuperHero, error) {
	query := `UPDATE super_heroes SET name=$1, updated_at=$2 WHERE id=$3;`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		superHero.Name,
		superHero.UpdatedAt,
		superHero.ID,
	)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return superHero, nil
}

func (this *postgresSuperHeroRepo) FindByUUID(ctx context.Context, superUUID int64) (*entities.SuperHero, error) {
	query := "SELECT id, uuid, name, created_at, updated_at FROM super_heroes WHERE uuid=$1"

	rows, err := this.fetch(ctx, query, superUUID)
	if err != nil {
		return nil, err
	}
	payload := &entities.SuperHero{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrSuperHeroNotFound
	}
	return payload, nil
}

func (this *postgresSuperHeroRepo) Delete(ctx context.Context, superUUID string) error {
	query := `DELETE FROM super_heroes WHERE uuid=$1`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, superUUID)
	if err != nil {
		return err
	}
	return nil
}

func (this *postgresSuperHeroRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.SuperHero, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.SuperHero, 0)
	for rows.Next() {
		superHero := new(entities.SuperHero)
		err := rows.Scan(
			&superHero.ID,
			&superHero.UUID,
			&superHero.Name,
			&superHero.CreatedAt,
			&superHero.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, superHero)
	}
	return payload, nil
}
