package factories

import "github.com/Thomasevano/EasyDocker/src/models"

func BuildErrorResponse(status string, message string) models.ErrorResponse {
	return models.ErrorResponse{
		Status:  status,
		Message: message,
	}
}
