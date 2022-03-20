package services

import (
	"github.com/cbodonnell/til-api/models"
	"github.com/cbodonnell/til-api/repositories"
)

type StandardTilService struct {
	repo repositories.TilRepository
}

func NewStandardTilService(_repo repositories.TilRepository) TilService {
	return &StandardTilService{
		repo: _repo,
	}
}

func (s *StandardTilService) GetAllByUserID(user_uuid string) ([]models.Til, error) {
	return s.repo.GetAllByUserID(user_uuid)
}

func (s *StandardTilService) Create(user_uuid string, til models.Til) (models.Til, error) {
	return s.repo.Create(user_uuid, til)
}
