package image

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CezarGarrido/careers/entities"
)

var ErrImageNotFound = errors.New("Image not found")

func NewPgSQLImageRepo(Conn *sql.DB) ImageRepo {
	return &postgresImageRepo{Conn: Conn}
}

type postgresImageRepo struct {
	Conn *sql.DB
}

func (this *postgresImageRepo) Create(ctx context.Context, image *entities.Image) (int64, error) {

	query := `INSERT INTO super_hero_images (uuid, super_hero_id, url, created_at) VALUES($1,$2,$3,$4) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		image.UUID,
		image.SuperID,
		image.URL,
		image.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresImageRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Image, error) {
	query := "SELECT id, uuid, super_hero_id, url, created_at, updated_at FROM super_hero_Image WHERE super_hero_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Image{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrImageNotFound
	}
	return payload, nil
}

func (this *postgresImageRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Image, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Image, 0)
	for rows.Next() {
		image := new(entities.Image)
		err := rows.Scan(
			&image.ID,
			&image.UUID,
			&image.SuperID,
			&image.URL,
			&image.CreatedAt,
			&image.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, image)
	}
	return payload, nil
}
