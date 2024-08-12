package links

import (
	"errors"
	"github.com/Braullio/linkito/internal/database"
	"github.com/Braullio/linkito/internal/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Create
// @Summary      Criação de encurtador de link
// @Tags         Links
// @Accept       json
// @Produce      json
// @Param        body body     links.Request  true  "JSON Body"
// @Success      201  {object} links.Link
// @Failure      400  {object} response.Error
// @Failure      500  {object} response.Error
// @Router       /v1/links [post]
func Create(c *fiber.Ctx) error {
	dto, err := biuldDTO(c)
	if err != nil {
		return response.ReturnError(c, fiber.StatusBadRequest, err.Error())
	}

	var link Link
	link.Build(dto)
	if err := link.Create(); err != nil {
		return response.ReturnError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}

// ListAll
// @Summary      Mostra uma lista de todos os links cadastrados
// @Tags         Links
// @Success      200  {array}   links.Link
// @Failure      500  {object}  response.Error
// @Router       /v1/links [get]
func ListAll(c *fiber.Ctx) error {
	var links []Link
	if err := database.GetDB().Find(&links).Error; err != nil {
		return response.ReturnError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(links)
}

// Search
// @Summary      Mostra os dados do link com base no ID fornecido
// @Tags         Links
// @Param        id   path      string  true  "Links ID"
// @Success      200  {object}  links.Link
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /v1/links/{id} [get]
func Search(c *fiber.Ctx) error {
	id := c.Params("id")
	var link Link
	if err := database.GetDB().Where("id = ?", id).First(&link).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return response.ReturnError(c, fiber.StatusNotFound, "Link not found")
		}
		return response.ReturnError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(link)
}

// Redirect
// @Summary      Redireciona para uma URL com base no ID fornecido
// @Tags         Links
// @Param        id   path      string  true  "Links ID"
// @Success      307
// @Failure      500  {object}  response.Error
// @Router       /{id} [get]
func Redirect(c *fiber.Ctx) error {
	id := c.Params("id")
	var link Link
	if err := database.GetDB().Where("id = ?", id).First(&link).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.ReturnError(c, fiber.StatusNotFound, "Link not found")
		}
		return response.ReturnError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Redirect(link.Link, fiber.StatusTemporaryRedirect)
}

func Migrate(c *fiber.Ctx) error {
	err := database.GetDB().AutoMigrate(Link{})
	if err != nil {
		return response.ReturnError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}
