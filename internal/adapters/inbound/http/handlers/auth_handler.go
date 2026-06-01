package handler

import (
	requestDTO "userAuth/internal/adapters/inbound/http/request_DTO"
	responseDTO "userAuth/internal/adapters/inbound/http/response"
	inboundPort "userAuth/internal/core/ports/inbound"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type AuthHandler struct {
	authService inboundPort.AuthService
	log         zerolog.Logger
}

func NewAuthHandler(authService inboundPort.AuthService, logger zerolog.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         logger,
	}
}

// free naming — not bound to any interface
func (h *AuthHandler) HandleLogin(c *fiber.Ctx) error {
	var req requestDTO.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		h.log.Err(err).Str("Email", req.Email).Msg("Login Failed")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.log.Err(err).Str("Email", req.Email).Msg("Login Failed")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	resp := responseDTO.LoginResponse{
		AccessToken: token.AccessToken,
		TokenType: token.TokenType,
		ExpiresIn: token.ExpiresIn,
	}

	h.log.Info().Str("email", req.Email).Msg("Login successful")

	return c.Status(fiber.StatusOK).JSON(resp)
}
