package handler

import (
	_ "assignment/docs"
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
	r.Get("/main", middleware.AuthRequired(), h.GetMain)
}

// HealthCheck godoc
// @Summary Check API health status
// @Description Returns a message indicating the API is working
// @Tags health
// @Produce json
// @Success 200 {object} domain.ResponseStatus "API is healthy"
// @Router /health [get]
func (h *UserHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(domain.ResponseStatus{
		Message: "API is working.",
	})
}

// Login godoc
// @Summary Login user and generate access token
// @Description Authenticate user by userId and PIN, then return token in cookie
// @Tags user
// @Accept json
// @Produce json
// @Param loginRequest body domain.LoginRequest true "Login Request Payload"
// @Success 200 "OK"
// @Failure 400 {object} domain.ResponseStatus "Invalid request body"
// @Failure 401 {object} domain.ResponseStatus "Invalid PIN"
// @Failure 500 {object} domain.ResponseStatus "Failed to generate token or internal error"
// @Router /login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseStatus{
			Message: "Invalid request body",
		})
	}

	user, err := h.userService.GetUserById(req.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseStatus{
			Message: err.Error(),
		})
	}

	if !helper.CheckPassword(req.Password, user.PinHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.ResponseStatus{
			Message: "Invalid PIN.",
		})
	}

	res, err := helper.GenerateJWT(user.UserId, user.Name, 15*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseStatus{
			Message: "Failed to generate token",
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

// Profile godoc
// @Summary Get user profile for validate token is valid
// @Description Get current user's profile from token claims
// @Tags user
// @Produce json
// @Success 200 {object} domain.UserResponse
// @Failure 401 {object} domain.ResponseStatus "Unauthorized"
// @Router /profile [get]
// @Security ApiCookieAuth
func (h *UserHandler) Profile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*domain.Claims)

	return c.Status(fiber.StatusOK).JSON(domain.UserResponse{
		UserID: claims.UserID,
		Name:   claims.Name,
	})
}

// GetMain godoc
// @Summary Get main user profile
// @Description Get user profile by user ID
// @Tags user
// @Produce json
// @Success 200 {object} domain.GetUserMain "User profile data"
// @Failure 500 {object} domain.ResponseStatus "Internal server error"
// @Router /main [get]
// @Security ApiCookieAuth
func (h *UserHandler) GetMain(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*domain.Claims)

	profile, err := h.userService.GetUserProfile(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//profile, err := h.userService.GetUserProfile("000018b0e1a211ef95a30242ac180002")
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": err.Error(),
	//	})
	//}

	profile.UserID = claims.UserID
	profile.Name = claims.Name

	return c.Status(fiber.StatusOK).JSON(profile)
}

// GetUser godoc
// @Summary Get user
// @Description Retrieve user data
// @Tags user
// @Produce json
// @Success 200 {array} domain.UserResponse "List of users"
// @Failure 500 {object} domain.ResponseStatus "Internal server error"
// @Router /user [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.userService.GetUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// Logout godoc
// @Summary Logout user by clearing access token cookie
// @Description Clears the access_token cookie to logout the user
// @Tags user
// @Success 200 "Logout successful"
// @Router /logout [post]
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
