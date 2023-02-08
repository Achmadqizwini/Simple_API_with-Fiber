package delivery

import (
	"be13/ca/features/mahasiswa"
	"be13/ca/utilss/helper"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MahasiswaDelivery struct {
	mahasiswaService mahasiswa.ServiceInterface
}

func New(service mahasiswa.ServiceInterface, e *fiber.App) {
	handler := &MahasiswaDelivery{
		mahasiswaService: service,
	}

	e.Post("/mahasiswa", handler.Create)
	e.Put("/mahasiswa/:id", handler.Update)
	e.Delete("/mahasiswa/:id", handler.Delete)
	e.Get("/mahasiswa/:id", handler.Read)
}

func (delivery *MahasiswaDelivery) Create(c *fiber.Ctx) error {
	userInput := mahasiswa.Core{}

	if errBind := c.BodyParser(&userInput); errBind != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error binding data" + errBind.Error(),
		})
	}

	if err := delivery.mahasiswaService.Create(userInput); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error create new data" + err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "succes",
		"message": "success create new users",
	})
}

func (delivery *MahasiswaDelivery) Update(c *fiber.Ctx) error {
	id, errConv := strconv.Atoi(c.Params("id"))
	if errConv != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":   "failed",
			"messsage": "failed convert param id" + errConv.Error(),
		})
	}

	userInput := mahasiswa.Core{}

	if errBind := c.BodyParser(&userInput); errBind != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error binding data" + errBind.Error(),
		})
	}

	if errUpt := delivery.mahasiswaService.Update(userInput, id); errUpt != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error Db update " + errUpt.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success update data",
	})
}

func (delivery *MahasiswaDelivery) Delete(c *fiber.Ctx) error {
	id, errConv := strconv.Atoi(c.Params("id"))
	if errConv != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.BadRequest(errConv.Error()))
	}

	if errDel := delivery.mahasiswaService.Delete(id); errDel != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.FailedResponse("error delete user" + errDel.Error()))
	}

	return c.Status(http.StatusOK).JSON(helper.SuccessResponse("success delete data"))
}

func (delivery *MahasiswaDelivery) Read(c *fiber.Ctx) error {
	id, errConv := strconv.Atoi(c.Params("id"))
	if errConv != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.BadRequest(errConv.Error()))
	}

	results, err := delivery.mahasiswaService.Read(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.FailedResponse(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(helper.SuccessWithDataResponse("Success read all data", results))
}
