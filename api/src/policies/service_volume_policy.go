package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

func CanAccessServiceVolume(user models.UserResponse, serviceVolumeId string) bool {
	serviceVolume, result := repositories.FindServiceRelation[models.ServiceVolume](serviceVolumeId)

	if result.RowsAffected == 0 {
		return false
	}

	return CanAccessService(user, serviceVolume.ServiceID)
}
