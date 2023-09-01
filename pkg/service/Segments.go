package service

import (
	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/Romon001/AvitoStart-app/pkg/repository"
)

type SegmentsService struct {
	repo repository.Segments
}

func NewSegmentService(repo repository.Segments) *SegmentsService {
	return &SegmentsService{repo: repo}
}

func (s *SegmentsService) Create(name string) (int, error) {
	return s.repo.Create(name)
}
func (s *SegmentsService) GetAll() ([]avitoStartApp.Segment, error) {
	return s.repo.GetAll()

}
func (s *SegmentsService) DeleteSegment(name string) error {
	return s.repo.DeleteSegment(name)
}
