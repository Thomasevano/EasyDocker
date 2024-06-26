package repositories

import (
	"github.com/Thomasevano/EasyDocker/src/initializers"
	"github.com/Thomasevano/EasyDocker/src/models"
	"gorm.io/gorm"
)

func FindServiceManagedVolumeLinksByStackId(stackId string) ([]models.ServiceManagedVolumeLink, *gorm.DB) {
	var links []models.ServiceManagedVolumeLink
	db := initializers.DB.
		Joins("Service").
		Where("stack_id = ?", stackId).
		Find(&links)

	return links, db
}

func FindServiceManagedVolumeLinkByServiceAndVolume(serviceId string, volumeId string) (models.ServiceManagedVolumeLink, *gorm.DB) {
	var serviceManagedVolumeLink models.ServiceManagedVolumeLink
	db := initializers.DB.
		Where("service_id = ? AND managed_volume_id = ?", serviceId, volumeId).
		First(&serviceManagedVolumeLink)

	return serviceManagedVolumeLink, db
}
