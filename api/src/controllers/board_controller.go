package controllers

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/policies"
	"github.com/Thomasevano/EasyDocker/src/repositories"
	"github.com/Thomasevano/EasyDocker/src/services/factories"
	"github.com/gofiber/fiber/v2"
)

// GetBoard godoc
// @Summary      Get board
// @Tags         Board
// @Accept       json
// @Produce      json
// @Param stackId path string true "Stack ID"
// @Success      200  {array}  models.Board
// @Router       /stacks/{stackId}/board [get]
func GetBoard(c *fiber.Ctx) error {
	stackId := c.Params("stackId")
	currentUser := c.Locals("user").(models.UserResponse)

	if !policies.CanAccessStack(currentUser, stackId) {
		return c.Status(fiber.StatusNotFound).JSON(factories.BuildErrorResponse("error", "Board not found"))
	}

	services, _ := repositories.FindServicesByStackIdWithAssociation(stackId)
	networks, _ := repositories.FindNetworksByStackId(stackId)
	volumes, _ := repositories.FindManagedVolumesByStackId(stackId)
	serviceNetworkLinks, _ := repositories.FindServiceNetworkLinksByStackId(stackId)
	serviceManagedVolumeLinks, _ := repositories.FindServiceManagedVolumeLinksByStackId(stackId)

	board := models.Board{
		Services:                  factories.BuildServiceResponses(services),
		Networks:                  factories.BuildNetworkBoardResponses(networks),
		Volumes:                   factories.BuildManagedVolumeBoardResponses(volumes),
		ServiceNetworkLinks:       factories.BuildServiceNetworkLinkResponses(serviceNetworkLinks),
		ServiceManagedVolumeLinks: factories.BuildServiceManagedVolumeLinkResponses(serviceManagedVolumeLinks),
	}

	return c.Status(fiber.StatusOK).JSON(board)
}
