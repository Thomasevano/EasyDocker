package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

func CanAccessServiceManagedVolumeLink(currentUser models.UserResponse, id string) bool {
	link, db := repositories.Find[models.ServiceManagedVolumeLink](id)
	var linkCount int64
	db.Count(&linkCount)

	if linkCount != 1 {
		return false
	}

	return CanAccessService(currentUser, link.ServiceID) &&
		CanAccessVolume(currentUser, link.ManagedVolumeID)
}
