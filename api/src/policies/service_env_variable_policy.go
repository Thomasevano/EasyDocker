package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

func CanAccessServiceEnvVariable(user models.UserResponse, serviceEnvVariableId string) bool {
	serviceEnvVariable, result := repositories.FindServiceRelation[models.ServiceEnvVariable](serviceEnvVariableId)

	if result.RowsAffected == 0 {
		return false
	}

	return CanAccessService(user, serviceEnvVariable.ServiceID)
}
