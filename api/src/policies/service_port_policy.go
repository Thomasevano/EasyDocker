package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

func CanAccessServicePort(user models.UserResponse, servicePortId string) bool {
	servicePort, result := repositories.FindServicePort(servicePortId)

	if result.RowsAffected == 0 {
		return false
	}

	return CanAccessService(user, servicePort.ServiceID)
}
