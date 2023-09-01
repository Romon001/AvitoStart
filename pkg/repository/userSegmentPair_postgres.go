package repository

import (
	"fmt"

	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/jmoiron/sqlx"
)

type UserSegmentPairPG struct {
	db *sqlx.DB
}

func newUserSegmentPairPostgres(db *sqlx.DB) *SegmentsPostgres {
	return &SegmentsPostgres{db: db}
}

func (r *SegmentsPostgres) AddUserToSegments(segmentList []avitoStartApp.Segment, userId int) (int, error) {
	var id int
	for i := 0; i < len(segmentList); i++ {
		query := fmt.Sprintf("INSERT INTO %s (segment,user) values ($1,$2) returning id", segmentTable)
		row := r.db.QueryRow(query, segmentList[i], userId)
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}

	return len(segmentList), nil
}

func (r *SegmentsPostgres) GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error) {
	var lists []avitoStartApp.UserSegmentPair

	query := fmt.Sprintf("SELECT name FROM %s where userId = $1", userSegmentPairtable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *SegmentsPostgres) DeleteUserFromSegments(segmentList []avitoStartApp.Segment, userId int) (int, error) {
	for i := 0; i < len(segmentList); i++ {
		query := fmt.Sprintf("DELETE FROM %s WHERE name = $1 and user = $2",
			userSegmentPairtable)
		_, err := r.db.Exec(query, segmentList[i], userId)
		if err != nil {
			return -1, err
		}
	}

	return len(segmentList), nil
}
