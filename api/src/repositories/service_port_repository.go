package repositories

import (
	"github.com/Thomasevano/EasyDocker/src/initializers"
	"github.com/Thomasevano/EasyDocker/src/models"
	"gorm.io/gorm"
)

func FindServicePort(id string) (models.ServicePort, *gorm.DB) {
	var servicePort models.ServicePort
	result := initializers.DB.First(&servicePort, "id = ?", id)
	return servicePort, result
}

func FindServicePortsByServiceId(serviceId string) ([]models.ServicePort, *gorm.DB) {
	var servicePorts []models.ServicePort
	result := initializers.DB.Where("service_id = ?", serviceId).Find(&servicePorts)
	return servicePorts, result
}
