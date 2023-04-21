package repository

import (
	"context"
	"github.com/andreibthoughtworks/challenges/solution/internal/domain"
)

type AdSQL struct {
	db SQL
}

func NewAdSQL(db SQL) AdSQL {
	return AdSQL{
		db: db,
	}
}

func (a AdSQL) Create(ctx context.Context, ad domain.Ad) (domain.Ad, error) {
	var query = `
		INSERT INTO 
			ads (id, title, description, price, created_at)
		VALUES 
			($1, $2, $3, $4, $5)
	`

	if err := a.db.ExecuteContext(
		ctx,
		query,
		ad.ID,
		ad.Title,
		ad.Description,
		ad.Price,
		ad.CreatedAt,
	); err != nil {
		return domain.Ad{}, err
	}

	return ad, nil
}
