package policies

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/repositories"
)

func CanAccessStack(user models.UserResponse, stackId string) bool {
	result, _ := repositories.GetStackByIdForAUser(stackId, user.ID)

	return result.RowsAffected > 0
}
