package connections

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CezarGarrido/careers/entities"
)

var ErrConnectionsNotFound = errors.New("Connections not found")

func NewPgSQLConnectionsRepo(Conn *sql.DB) ConnectionsRepo {
	return &postgresConnectionsRepo{Conn: Conn}
}

type postgresConnectionsRepo struct {
	Conn *sql.DB
}

func (this *postgresConnectionsRepo) Create(ctx context.Context, connections *entities.Connections) (int64, error) {

	query := `INSERT INTO super_hero_connections (uuid, super_hero_id, group_affiliation, relatives, created_at) VALUES($1,$2,$3,$4,$5)RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		connections.UUID,
		connections.SuperID,
		connections.GroupAffiliation,
		connections.Relatives,
		connections.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresConnectionsRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Connections, error) {
	query := "SELECT id, uuid, super_hero_id, gender, race, height, weight, eye_color, hair_color, created_at, updated_at FROM super_hero_Connections WHERE super_hero_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Connections{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrConnectionsNotFound
	}
	return payload, nil
}

func (this *postgresConnectionsRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Connections, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Connections, 0)
	for rows.Next() {
		connections := new(entities.Connections)
		err := rows.Scan(
			&connections.ID,
			&connections.UUID,
			&connections.SuperID,
			&connections.GroupAffiliation,
			&connections.Relatives,
			&connections.CreatedAt,
			&connections.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, connections)
	}
	return payload, nil
}
