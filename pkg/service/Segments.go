package service

import "github.com/Romon001/AvitoStart-app/pkg/repository"

type SegmentsService struct {
	repo repository.Segments
}

func NewSegmentService(repo repository.Segments) *SegmentsService {
	return &SegmentsService{repo: repo}
}

func (s *SegmentsService) Create(name string) (int, error) {
	return s.repo.Create(name)
}
