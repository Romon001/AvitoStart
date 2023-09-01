package repository

import (
	"fmt"

	avitoStartApp "github.com/Romon001/AvitoStart-app"
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
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) returning id", segmentTable)
	row := r.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SegmentsPostgres) GetAll() ([]avitoStartApp.Segment, error) {
	var lists []avitoStartApp.Segment

	query := fmt.Sprintf("SELECT name FROM %s", segmentTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *SegmentsPostgres) DeleteSegment(name string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1",
		segmentTable)
	_, err := r.db.Exec(query, name)

	return err
}
