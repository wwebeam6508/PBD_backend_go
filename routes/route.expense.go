package routes

import (
	controller "PBD_backend_go/controller/expense"
	"PBD_backend_go/middleware"
	"PBD_backend_go/model"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoute(route fiber.Router) {
	route.Get("/get", middleware.Authenication, func(c *fiber.Ctx) error {
		return middleware.Permission(c, model.PermissionInput{
			GroupName: "Expense",
			Name:      "CanView",
		})
	}, controller.GetExpenseController)
}
