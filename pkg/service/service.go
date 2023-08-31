package service

import (
	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/Romon001/AvitoStart-app/pkg/repository"
)

type Service struct {
	Segments
	UserSegmentPair
}
type Segments interface {
	Create(name string) (int, error)
	GetAll(id int) ([]avitoStartApp.Segment, error)
	DeleteSegment(name string) error
}

type UserSegmentPair interface {
	AddUserToSegments(segmentList []string, userId int) (int, error)
	GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error)
	DeleteUserFromSegments(segmentList []string, userId int) (int, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
