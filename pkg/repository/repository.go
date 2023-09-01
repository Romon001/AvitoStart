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
	DeleteSegment(name string) error
}

type UserSegmentPair interface {
	AddUserToSegments(segmentList []avitoStartApp.Segment, userId int) (int, error)
	GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error)
	DeleteUserFromSegments(segmentList []avitoStartApp.Segment, userId int) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segments:        newSegmentsPostgres(db),
		UserSegmentPair: newUserSegmentPairPostgres(db),
	}
}
