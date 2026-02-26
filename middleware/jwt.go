package middleware

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

func GenerateJWT(email string, user uint, role string) (string, string) {

	tokenA, err := jwt.NewBuilder().
		Issuer("clubinformatico").
		IssuedAt(time.Now()).
		Subject("juackomdz").
		Claim("user", user).
		Claim("email", email).
		Claim("role", role).
		Expiration(time.Now().Add(5 * time.Minute)).
		Build()

	if err != nil {
		log.Print(err)
	}

	signedA, errort := jwt.Sign(tokenA, jwt.WithKey(jwa.HS256(), []byte(os.Getenv("JWT_SECRET"))), jwt.WithBase64Encoder(base64.URLEncoding))
	if errort != nil {
		log.Print(errort)
	}

	tokenR, errR := jwt.NewBuilder().
		Issuer("clubinformatico").
		IssuedAt(time.Now()).
		Subject("juackomdz").
		Claim("user", user).
		Claim("email", email).
		Claim("role", role).
		Expiration(time.Now().Add(24 * time.Hour)).
		Build()

	if errR != nil {
		log.Print(errR)
	}

	signedR, errorR := jwt.Sign(tokenR, jwt.WithKey(jwa.HS256(), []byte(os.Getenv("JWT_SECRET"))), jwt.WithBase64Encoder(base64.URLEncoding))
	if errorR != nil {
		log.Print(errorR)
	}

	return string(signedA), string(signedR)
}

func RefreshToken(token string) (string, error) {

	parsed, errV := jwt.Parse([]byte(token), jwt.WithKey(jwa.HS256(), []byte(os.Getenv("JWT_SECRET"))), jwt.WithValidate(true), jwt.WithBase64Encoder(base64.URLEncoding))
	if errV != nil {
		log.Print(errV)
	}

	var user float64
	var email string
	var role string

	parsed.Get("user", &user)
	parsed.Get("email", &email)
	parsed.Get("role", &role)

	refresh, err := jwt.NewBuilder().
		Issuer("clubinformatico").
		IssuedAt(time.Now()).
		Subject("juackomdz").
		Claim("user", uint(user)).
		Claim("email", email).
		Claim("role", role).
		Expiration(time.Now().Add(15 * time.Minute)).
		Build()

	if err != nil {
		return "", err
	}

	signedRefresh, errort := jwt.Sign(refresh, jwt.WithKey(jwa.HS256(), []byte(os.Getenv("JWT_SECRET"))), jwt.WithBase64Encoder(base64.URLEncoding))
	if errort != nil {
		return "", errort
	}

	return string(signedRefresh), nil
}

func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			role := c.Get("role")
			for _, r := range allowedRoles {
				if role == r {
					return next(c)
				}
			}

			return c.JSON(403, echo.Map{"mensaje": "No tienes permisos para este recurso"})
		}
	}
}

func MiddleJWT() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			h := c.Request().Header.Get("Authorization")
			if len(h) == 0 {
				return c.JSON(400, echo.Map{"mensaje": "Falta token de autorizacion"})
			}

			strToken := strings.TrimSpace(strings.Replace(h, "Bearer", "", 1))

			parsed, errp := jwt.Parse([]byte(strToken), jwt.WithKey(jwa.HS256(), []byte(os.Getenv("JWT_SECRET"))), jwt.WithValidate(true), jwt.WithBase64Encoder(base64.URLEncoding))
			if errp != nil {
				log.Print(errp)
				return c.JSON(401, echo.Map{"mensaje": "Error con token, generelo nuevamente"})
			}

			var role string
			var user float64

			parsed.Get("role", &role)
			parsed.Get("user", &user)

			c.Set("role", role)
			c.Set("user", uint(user))

			return next(c)
		}
	}
}
