package controllers

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/policies"
	"github.com/Thomasevano/EasyDocker/src/repositories"
	"github.com/Thomasevano/EasyDocker/src/services/docker_compose"
	"github.com/gofiber/fiber/v2"
)

func GenerateDockerComposeFile(c *fiber.Ctx) error {
	stackId := c.Params("stackId")
	currentUser := c.Locals("user").(models.UserResponse)

	if !policies.CanAccessStack(currentUser, stackId) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Stack not found"})
	}

	services, _ := repositories.FindServicesByStackIdWithAssociation(stackId)
	networks, _ := repositories.FindNetworksByStackId(stackId)
	volumes, _ := repositories.FindManagedVolumesByStackId(stackId)

	yaml := docker_compose.GenerateDockerCompose(services, networks, volumes)

	return c.Status(fiber.StatusOK).Send([]byte(yaml))
}
