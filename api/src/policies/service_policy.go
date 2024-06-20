package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

// CanAccessService checks if a user can access a service by
// searching if the service exists and if the user has access to the stack the service belongs to
func CanAccessService(user models.UserResponse, serviceId string) bool {
	service, result := repositories.FindService(serviceId)

	if result.RowsAffected == 0 {
		return false
	}

	return CanAccessStack(user, service.StackID)
}
