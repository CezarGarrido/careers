package work

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CezarGarrido/careers/entities"
)

var ErrWorkNotFound = errors.New("Work not found")

func NewPgSQLWorkRepo(Conn *sql.DB) WorkRepo {
	return &postgresWorkRepo{Conn: Conn}
}

type postgresWorkRepo struct {
	Conn *sql.DB
}

func (this *postgresWorkRepo) Create(ctx context.Context, work *entities.Work) (int64, error) {

	query := `INSERT INTO super_hero_works (uuid, super_hero_id, occupation, base, created_at) VALUES($1,$2,$3,$4,$5) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		work.UUID,
		work.SuperID,
		work.Occupation,
		work.BaseWork,
		work.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresWorkRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Work, error) {
	query := "SELECT id, uuid, super_hero_id, occupation, base, created_at, updated_at FROM super_hero_work WHERE super_hero_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Work{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrWorkNotFound
	}
	return payload, nil
}

func (this *postgresWorkRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Work, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Work, 0)
	for rows.Next() {
		work := new(entities.Work)
		err := rows.Scan(
			&work.ID,
			&work.UUID,
			&work.SuperID,
			&work.Occupation,
			&work.BaseWork,
			&work.CreatedAt,
			&work.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, work)
	}
	return payload, nil
}
