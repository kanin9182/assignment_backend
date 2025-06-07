package handler

import (
	"assignment/internals/core/domain"
	"assignment/internals/core/ports/services"
	"assignment/internals/helper"
	"assignment/internals/middleware"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserHandler struct {
	userService services.UserServices
}

func NewUserHandler(userService services.UserServices) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/user", h.GetUser)
	r.Get("/health", h.HealthCheck)
	r.Post("/login", h.Login)
	r.Get("/profile", middleware.AuthRequired(), h.Profile)
	r.Post("/logout", middleware.AuthRequired(), h.Logout)
	r.Get("/main", h.GetMain)
}

func (h *UserHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "API is working",
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.GetUserById(req.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !helper.CheckPassword(req.Password, user.PinHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid PIN",
		})
	}

	res, err := helper.GenerateJWT(user.UserId, user.Name, 15*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/",
	})

	//return c.Status(fiber.StatusOK).JSON(&domain.LoginResponse{
	//	AccessToken: res,
	//	TokenType:   "Bearer",
	//	ExpiresIn:   int((60 * time.Minute).Seconds()),
	//	User: &domain.UserResponse{
	//		UserID: user.UserId,
	//		Name:   user.Name,
	//	},
	//})

	return c.SendStatus(fiber.StatusOK)
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*domain.Claims)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": claims.UserID,
		"name":    claims.Name,
	})
}

func (h *UserHandler) GetMain(c *fiber.Ctx) error {
	// claims := c.Locals("claims").(*domain.Claims)

	//profile, err := h.userService.GetUserProfile(claims.UserID)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": err.Error(),
	//	})
	//}

	profile, err := h.userService.GetUserProfile("000018b0e1a211ef95a30242ac180002")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//profile.UserID = claims.UserID
	//profile.Name = claims.Name

	return c.Status(fiber.StatusOK).JSON(profile)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.userService.GetUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/",
	})
	return c.SendStatus(fiber.StatusOK)
}
