package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("JWT AUTH BRO")
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		fmt.Println("token not present in the header")
		return fmt.Errorf("no auth bro")
	}
	claims, err := validateToken(strings.Join(token, ""))
	if err != nil {
		return err
	}
	expiresFloat := claims["expires"].(float64)
	expires := int64(expiresFloat)
	if time.Now().Unix() > expires {
		return fmt.Errorf("token expired bro")
	}
	fmt.Println(claims)
	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("INVALID SIGNING METHOD", token.Header["alg"])
			return nil, fmt.Errorf("no auth broski")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("never print secret aman", secret)
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse jwt token:", err)
		return nil, fmt.Errorf("NO AUTH")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")

	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
