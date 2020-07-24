package powerstats

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CezarGarrido/careers/entities"
)

var ErrPowerstatsNotFound = errors.New("Super hero not found")

func NewPgSQLPowerstatsRepo(Conn *sql.DB) PowerstatsRepo {
	return &postgresPowerstatsRepo{Conn: Conn}
}

type postgresPowerstatsRepo struct {
	Conn *sql.DB
}

func (this *postgresPowerstatsRepo) Create(ctx context.Context, powerstats *entities.Powerstats) (int64, error) {

	query := `INSERT INTO super_heroe_powerstats (uuid, super_id, intelligence, strength, speed, durability, power, combat, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		powerstats.UUID,
		powerstats.SuperID,
		powerstats.Intelligence,
		powerstats.Strength,
		powerstats.Speed,
		powerstats.Durability,
		powerstats.Power,
		powerstats.Combat,
		powerstats.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresPowerstatsRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Powerstats, error) {
	query := "SELECT id, uuid, super_id, intelligence, strength, speed, durability, power, combat, created_at, updated_at FROM super_heroe_powerstats WHERE super_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Powerstats{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrPowerstatsNotFound
	}
	return payload, nil
}

func (this *postgresPowerstatsRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Powerstats, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Powerstats, 0)
	for rows.Next() {
		powerstats := new(entities.Powerstats)
		err := rows.Scan(
			&powerstats.ID,
			&powerstats.UUID,
			&powerstats.SuperID,
			&powerstats.Intelligence,
			&powerstats.Strength,
			&powerstats.Speed,
			&powerstats.Durability,
			&powerstats.Power,
			&powerstats.Combat,
			&powerstats.CreatedAt,
			&powerstats.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, powerstats)
	}
	return payload, nil
}
