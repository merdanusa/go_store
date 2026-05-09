package utils

import "github.com/gofiber/fiber/v3"

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

func SendError(c fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(ErrorResponse{
		Success: false,
		Error:   message,
		Code:    status,
		Details: details,
	})
}

func SendSuccess(c fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(SuccessResponse{
		Success: true,
		Message: message,
		Code:    status,
		Details: details,
	})
}
