package repositories

import (
	"github.com/Thomasevano/EasyDocker/src/initializers"
	"github.com/Thomasevano/EasyDocker/src/models"
	"gorm.io/gorm"
)

func FindServiceNetworkLinksByStackId(stackId string) ([]models.ServiceNetworkLink, *gorm.DB) {
	var serviceNetworkLinks []models.ServiceNetworkLink
	db := initializers.DB.
		Joins("Service").
		Where("stack_id = ?", stackId).
		Find(&serviceNetworkLinks)

	return serviceNetworkLinks, db
}

func FindServiceNetworkLinkByServiceAndNetwork(serviceId string, networkId string) (models.ServiceNetworkLink, *gorm.DB) {
	var serviceNetworkLink models.ServiceNetworkLink
	db := initializers.DB.
		Where("service_id = ? AND network_id = ?", serviceId, networkId).
		First(&serviceNetworkLink)

	return serviceNetworkLink, db
}

func FindServiceNetworkLink(id string) (models.ServiceNetworkLink, *gorm.DB) {
	var serviceNetworkLink models.ServiceNetworkLink
	db := initializers.DB.
		Where("id = ?", id).
		First(&serviceNetworkLink)

	return serviceNetworkLink, db
}
