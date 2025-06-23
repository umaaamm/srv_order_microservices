package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"srv_order/main/api/presenter"
	"srv_order/main/pkg/entities"
	contact "srv_order/main/pkg/order"
)

func AddOrder(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Order
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		if requestBody.Nama == "" || requestBody.NoHp == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(errors.New(
				"Please specify name and noHp")))
		}

		resultFromContact, err := service.GetContactFromGRPC(requestBody.NoHp)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}

		requestBody.Order = resultFromContact.Nama // for sample only

		result, err := service.InsertOrder(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(presenter.ContactSuccessResponse(result))
	}
}

func Gets(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchOrders()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(presenter.ContactsSuccessResponse(fetched))
	}
}
