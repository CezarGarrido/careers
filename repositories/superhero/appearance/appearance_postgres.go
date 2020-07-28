package appearance

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/CezarGarrido/careers/entities"
)

var ErrAppearanceNotFound = errors.New("Appearance not found")

func NewPgSQLAppearanceRepo(Conn *sql.DB) AppearanceRepo {
	return &postgresAppearanceRepo{Conn: Conn}
}

type postgresAppearanceRepo struct {
	Conn *sql.DB
}

func (this *postgresAppearanceRepo) Create(ctx context.Context, appearance *entities.Appearance) (int64, error) {

	query := `INSERT INTO super_hero_appearance (uuid, super_hero_id, gender, race, height, weight, eye_color, hair_color, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	var returnID int64

	err = stmt.QueryRowContext(ctx,
		appearance.UUID,
		appearance.SuperID,
		appearance.Gender,
		appearance.Race,
		strings.Join(appearance.Height, ","),
		strings.Join(appearance.Weight, ","),
		appearance.EyeColor,
		appearance.HairColor,
		appearance.CreatedAt,
	).Scan(&returnID)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return returnID, nil
}

func (this *postgresAppearanceRepo) FindBySuperID(ctx context.Context, superID int64) (*entities.Appearance, error) {
	query := "SELECT id, uuid, super_hero_id, gender, race, height, weight, eye_color, hair_color, created_at updated_at FROM super_hero_appearance WHERE super_hero_id=$1"

	rows, err := this.fetch(ctx, query, superID)
	if err != nil {
		return nil, err
	}
	payload := &entities.Appearance{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, ErrAppearanceNotFound
	}
	return payload, nil
}

func (this *postgresAppearanceRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Appearance, error) {

	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Appearance, 0)
	for rows.Next() {
		appearance := new(entities.Appearance)
		err := rows.Scan(
			&appearance.ID,
			&appearance.UUID,
			&appearance.SuperID,
			&appearance.Gender,
			&appearance.Race,
			&appearance.Height,
			&appearance.Weight,
			&appearance.EyeColor,
			&appearance.HairColor,
			&appearance.CreatedAt,
			&appearance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, appearance)
	}
	return payload, nil
}
