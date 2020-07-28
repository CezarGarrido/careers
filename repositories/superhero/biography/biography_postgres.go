package biography

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/CezarGarrido/careers/entities"
)

var ErrBiographyNotFound = errors.New("Biography not found")

func NewPgSQLBiographyRepo(Conn *sql.DB) BiographyRepo {
	return &postgresBiographyRepo{Conn: Conn}
}

type postgresBiographyRepo struct {
	Conn *sql.DB
}

func (this *postgresBiographyRepo) Create(ctx context.Context, biography *entities.Biography) (int64, error) {

	query := `INSERT INTO super_hero_biography (uuid, super_hero_id, fullname, alter_egos, aliases, place_of_birth, first_appearance, publisher, alignment, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		biography.UUID,
		biography.SuperID,
		biography.FullName,
		biography.AlterEgos,
		strings.Join(biography.Aliases, ","),
		biography.PlaceOfBirth,
		biography.FirstAppearance,
		biography.Publisher,
		biography.Alignment,
		biography.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresBiographyRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Biography, error) {
	query := "SELECT id, uuid, super_hero_id, fullname, alter_egos, aliases, place_of_birth, first_appearance, publisher, alignment, created_at, updated_at FROM super_hero_biography WHERE super_hero_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Biography{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrBiographyNotFound
	}
	return payload, nil
}

func (this *postgresBiographyRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Biography, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Biography, 0)
	for rows.Next() {
		biography := new(entities.Biography)
		err := rows.Scan(
			&biography.ID,
			&biography.UUID,
			&biography.SuperID,
			&biography.FullName,
			&biography.AlterEgos,
			&biography.Aliases,
			&biography.PlaceOfBirth,
			&biography.FirstAppearance,
			&biography.Publisher,
			&biography.Alignment,
			&biography.CreatedAt,
			&biography.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, biography)
	}
	return payload, nil
}
