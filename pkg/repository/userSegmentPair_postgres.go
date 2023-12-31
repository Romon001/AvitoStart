package repository

import (
	"fmt"

	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/jmoiron/sqlx"
)

type UserSegmentPairPG struct {
	db *sqlx.DB
}

func newUserSegmentPairPostgres(db *sqlx.DB) *UserSegmentPairPG {
	return &UserSegmentPairPG{db: db}
}

func (r *UserSegmentPairPG) AddUserToSegments(segmentList []string, userId int) (int, error) {
	var id int
	for i := 0; i < len(segmentList); i++ {
		query := fmt.Sprintf("INSERT INTO %s (userid,segmentname) values ($2, $1) returning id", userSegmentPairtable)
		row := r.db.QueryRow(query, segmentList[i], userId)
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}

	return len(segmentList), nil
}

func (r *UserSegmentPairPG) GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error) {
	var lists []avitoStartApp.UserSegmentPair

	query := fmt.Sprintf("SELECT segmentname, userid FROM %s where userid = $1", userSegmentPairtable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *UserSegmentPairPG) DeleteUserFromSegments(segmentList []string, userId int) (int, error) {
	for i := 0; i < len(segmentList); i++ {
		query := fmt.Sprintf("DELETE FROM %s WHERE segmentname = $1 and userid = $2",
			userSegmentPairtable)
		_, err := r.db.Exec(query, segmentList[i], userId)
		if err != nil {
			return -1, err
		}
	}

	return len(segmentList), nil
}
