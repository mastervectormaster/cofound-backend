package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/mastervectormaster/cofound-backend/utils"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	jwtMiddleware := jwtware.New(
		jwtware.Config{
			SigningKey: utils.JWTSecret,
			AuthScheme: "Token",
		})
	//(*v1).Use(jwtMiddleware)
	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
	guestUsers.Post("/login", h.Login)
	user := v1.Group("/user", jwtMiddleware)
	user.Get("", h.CurrentUser)
	user.Put("", h.UpdateUser)
}
