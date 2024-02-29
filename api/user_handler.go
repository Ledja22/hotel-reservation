package api

import (
	"github.com/Ledja22/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Call",
	}
	return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("james only by id")
}
