package service

import (
	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/Romon001/AvitoStart-app/pkg/repository"
)

type UserSegmentPairService struct {
	repo repository.UserSegmentPair
}

func NewUserSegmentPairService(repo repository.UserSegmentPair) *UserSegmentPairService {
	return &UserSegmentPairService{repo: repo}
}

func (s *UserSegmentPairService) AddUserToSegments(segmentList []avitoStartApp.Segment, userId int) (int, error) {
	return s.repo.AddUserToSegments(segmentList, userId)
}
func (s *UserSegmentPairService) GetUserSegments(userId int) ([]avitoStartApp.UserSegmentPair, error) {
	return s.repo.GetUserSegments(userId)

}
func (s *UserSegmentPairService) DeleteUserFromSegments(segmentList []avitoStartApp.Segment, userId int) (int, error) {
	return s.repo.DeleteUserFromSegments(segmentList, userId)
}
