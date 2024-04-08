package middleware

import (
	"github.com/gofiber/fiber/v2"

	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func JwtProtect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header value
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Pesan": "Masukan JWT TOKEN",
			})
		}
		// Check if Authorization header value is in the correct format
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			return c.JSON(fiber.Map{
				"status": "Invalid",
			})
		}

		// Extract token string
		tokenString := authHeader[7:]

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method and secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("Rahasia"), nil
		})

		if err != nil {
			// Token not valid, handle error
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"Pesan": "Maaf Gagal Mengambil Token",
				})
		}

		// Get claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			// Claims not valid, handle error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Extract data from claims and add it to the context
		//c.Locals("name", claims["name"].(string))
		c.Locals("id_user", int(claims["id_user"].(float64)))
		c.Locals("role", claims["role"].(float64))
		//c.Locals("satminkal", claims["satminkal"].(string))

		// Continue with next middleware/handler
		return c.Next()
	}
}

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
