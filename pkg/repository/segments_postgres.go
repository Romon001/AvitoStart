package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentsPostgres struct {
	db *sqlx.DB
}

func newSegmentsPostgres(db *sqlx.DB) *SegmentsPostgres {
	return &SegmentsPostgres{db: db}
}

func (r *SegmentsPostgres) Create(name string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) returning id", "segments")
	row := r.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
