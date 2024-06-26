package controllers

import (
	"github.com/Thomasevano/EasyDocker/src/helpers"
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/policies"
	"github.com/Thomasevano/EasyDocker/src/repositories"
	"github.com/Thomasevano/EasyDocker/src/services/factories"
	"github.com/gofiber/fiber/v2"
)

// GetManagedVolume godoc
// @Summary      Get a volume
// @Tags         Managed Volumes
// @Accept       json
// @Produce      json
// @Param id path string true "Volume ID"
// @Success      200  {object}  models.ManagedVolumeResponse
// @Router       /managed_volumes/{id} [get]
func GetManagedVolume(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUser := c.Locals("user").(models.UserResponse)

	if !policies.CanAccessVolume(currentUser, id) {
		return c.Status(fiber.StatusNotFound).JSON(factories.BuildErrorResponse("error", "Volume not found"))
	}

	volume, _ := repositories.Find[models.ManagedVolume](id)

	return c.Status(fiber.StatusOK).JSON(factories.BuildManagedVolumeResponse(volume))
}

// CreateManagedVolume godoc
// @Summary      Create a volume
// @Tags         Managed Volumes
// @Accept       json
// @Produce      json
// @Param stackId path string true "Stack ID"
// @Param volume body models.ManagedVolumeCreateInput true "Volume"
// @Success      201  {object}  models.ManagedVolumeResponse
// @Router       /stacks/{stackId}/managed_volumes [post]
func CreateManagedVolume(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.UserResponse)
	stackId := c.Params("stackId")

	if !policies.CanAccessStack(currentUser, stackId) {
		return c.Status(fiber.StatusNotFound).JSON(factories.BuildErrorResponse("error", "Stack TOTO not found"))
	}

	body, err := helpers.BodyParse[models.ManagedVolumeCreateInput](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(factories.BuildErrorResponse("error", "Cannot parse JSON"))
	}

	volume := factories.BuildManagedVolumeFromCreationInput(body, stackId)

	result := repositories.Create[models.ManagedVolume](&volume)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(factories.BuildErrorResponse("error", "Cannot create service"))
	}

	return c.Status(fiber.StatusCreated).JSON(factories.BuildManagedVolumeResponse(volume))
}

// UpdateManagedVolume godoc
// @Summary      Update a volume
// @Tags         Managed Volumes
// @Accept       json
// @Produce      json
// @Param id path string true "Managed Volume ID"
// @Param request body models.ManagedVolumeUpdateInput true "query params"
// @Success      200  {object}  models.ManagedVolumeResponse
// @Router       /managed_volumes/{id} [put]
func UpdateManagedVolume(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.UserResponse)
	id := c.Params("id")

	if !policies.CanAccessVolume(currentUser, id) {
		return c.Status(fiber.StatusNotFound).JSON(factories.BuildErrorResponse("error", "Managed Volume not found"))
	}

	volume, _ := repositories.Find[models.ManagedVolume](id)

	body, err := helpers.BodyParse[models.ManagedVolumeUpdateInput](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(factories.BuildErrorResponse("error", "Cannot parse JSON"))
	}

	updatedVolume := factories.BuildManagedVolumeFromUpdateInput(body)

	result := repositories.Update(&volume, &updatedVolume)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(factories.BuildErrorResponse("error", "Cannot update network"))
	}

	return c.Status(fiber.StatusOK).JSON(factories.BuildManagedVolumeResponse(volume))
}

// DeleteManagedVolume godoc
// @Summary      Delete a volume
// @Tags         Managed Volumes
// @Accept       json
// @Produce      json
// @Param id path string true "Volume ID"
// @Success      204
// @Router       /managed_volumes/{id} [delete]
func DeleteManagedVolume(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUser := c.Locals("user").(models.UserResponse)

	if !policies.CanAccessVolume(currentUser, id) {
		return c.Status(fiber.StatusNotFound).JSON(factories.BuildErrorResponse("error", "Volume not found"))
	}

	repositories.DeleteById[models.ManagedVolume](id)

	return c.SendStatus(fiber.StatusNoContent)
}
