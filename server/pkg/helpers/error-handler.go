package helpers

import "github.com/gofiber/fiber/v2"

// Default error handler
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return statuscode with error message
	return c.Status(code).JSON(map[string]interface{}{
		"message": err.Error(),
	})
}
