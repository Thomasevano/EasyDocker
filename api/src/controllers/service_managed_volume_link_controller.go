package controllers

import (
	"github.com/Thomasevano/EasyDocker/src/helpers"
	"github.com/Thomasevano/EasyDocker/src/initializers"
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/policies"
	"github.com/Thomasevano/EasyDocker/src/repositories"
	"github.com/Thomasevano/EasyDocker/src/services/factories"
	"github.com/gofiber/fiber/v2"
)

// CreateServiceManagedVolumeLink godoc
// @Summary      Create a new link between a service and a volume
// @Tags         Service Volume Links
// @Accept       json
// @Produce      json
// @Param request body models.ServiceManagedVolumeLinkCreateInput true "query params"
// @Success      200  {object}  models.ServiceNetworkLinkResponse
// @Router       /service_managed_volume_links [post]
func CreateServiceManagedVolumeLink(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.UserResponse)

	body, err := helpers.BodyParse[models.ServiceManagedVolumeLinkCreateInput](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(factories.BuildErrorResponse("error", err.Error()))
	}

	errors := helpers.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if !policies.CanAccessService(currentUser, body.ServiceID) ||
		!policies.CanAccessVolume(currentUser, body.ManagedVolumeID) {
		return c.Status(fiber.StatusNotFound).
			JSON(factories.BuildErrorResponse("error", "Service or network not found"))
	}

	_, db := repositories.FindServiceManagedVolumeLinkByServiceAndVolume(body.ServiceID, body.ManagedVolumeID)

	if db.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).
			JSON(factories.BuildErrorResponse("error", "Link already exists"))
	}

	newServiceManagedVolumeLink := factories.BuildServiceManagedVolumeLinkFromCreateInput(body)

	result := initializers.DB.Create(&newServiceManagedVolumeLink)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).
			JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).
		JSON(factories.BuildServiceManagedVolumeLinkResponse(newServiceManagedVolumeLink))
}

// UpdateServiceManagedVolumeLink godoc
// @Summary      Update a link between a service and a volume
// @Tags         Service Volume Links
// @Accept       json
// @Produce      json
// @Param request body models.ServiceManagedVolumeLinkUpdateInput true "query params"
// @Param id path string true "Service Volume Link ID"
// @Success      200  {object}  models.ServiceNetworkLinkResponse
// @Router       /service_managed_volume_links/{id} [put]
func UpdateServiceManagedVolumeLink(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.UserResponse)
	linkId := c.Params("id")

	body, err := helpers.BodyParse[models.ServiceManagedVolumeLinkUpdateInput](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(factories.BuildErrorResponse("error", err.Error()))
	}

	errors := helpers.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if !policies.CanAccessServiceManagedVolumeLink(currentUser, linkId) {
		return c.Status(fiber.StatusNotFound).
			JSON(factories.BuildErrorResponse("error", "Service or volume not found"))
	}

	link, db := repositories.Find[models.ServiceManagedVolumeLink](linkId)

	if db.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).
			JSON(factories.BuildErrorResponse("error", "Link not found"))
	}

	updatedLink := factories.BuildServiceManagedVolumeLinkFromUpdateInput(body)

	result := repositories.Update(&link, &updatedLink)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).
			JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).
		JSON(factories.BuildServiceManagedVolumeLinkResponse(link))
}

// DeleteServiceManagedVolumeLink godoc
// @Summary      Delete a link between a service and a volume
// @Tags         Service Volume Links
// @Accept       json
// @Produce      json
// @Param id path string true "Service Volume Link ID"
// @Success      204
// @Router       /service_managed_volume_links/{id} [delete]
func DeleteServiceManagedVolumeLink(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.UserResponse)
	linkId := c.Params("id")

	if !policies.CanAccessServiceNetworkLink(currentUser, linkId) {
		return c.Status(fiber.StatusNotFound).
			JSON(factories.BuildErrorResponse("error", "Service or volume not found"))
	}

	link, db := repositories.FindServiceNetworkLink(linkId)

	result := db.Delete(&link)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).
			JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
