package repository

import (
	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Segments
	UserSegmentPair
}
type Segments interface {
	Create(name string) (int, error)
	GetAll() ([]avitoStartApp.Segment, error)
	//Delete(name string) error
}

type UserSegmentPair interface {
	AddUserToSegments(segmentList []string, userId int) (int, error)
	GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error)
	DeleteUserFromSegments(segmentList []string, userId int) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segments: newSegmentsPostgres(db),
	}
}
