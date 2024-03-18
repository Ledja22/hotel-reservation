package api

import (
	"github.com/Ledja22/hotel-reservation/db"
	"github.com/Ledja22/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {

	id := c.Params("id")
	user, err := h.userStore.GetUserById(id)
	if err != nil {

	}

	return c.JSON(u)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the watercooler",
	}
	return c.JSON(u)
}
