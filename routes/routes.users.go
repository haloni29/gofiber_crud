package routes

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func UserRouter(router fiber.Router) {

	Users := []User{
		{
			Id:       1,
			Username: "nicolas",
		},
		{
			Id:       2,
			Username: "Alejandro",
		},
	}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Users": Users,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		type request struct {
			Id       int
			Username string
		}
		var body request
		c.BodyParser(&body)

		NewUser := User{
			Id:       len(Users) + 1,
			Username: body.Username,
		}

		Users = append(Users, NewUser)

		return c.JSON(fiber.Map{
			"Users": Users,
		})
	})
	router.Delete("/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("Id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ERROR": "you have a error",
			})
		}

		for i, User := range Users {
			if User.Id == id {
				Users = append(Users[:i], Users[i+1:]...)
			}
		}

		return c.JSON(fiber.Map{
			"Users": Users,
		})
	})

}
