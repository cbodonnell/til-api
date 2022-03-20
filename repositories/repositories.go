package repositories

import (
	"github.com/cbodonnell/til-api/models"
)

type TilRepository interface {
	Close()
	GetAllByUserID(user_uuid string) ([]models.Til, error)
	Create(user_uuid string, til models.Til) (models.Til, error)
}
