package middleware

import (
	"encoding/base64"
	"log"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func GeneraJWT(email string, user uint16) string {

	token, err := jwt.NewBuilder().
		Issuer("clubinformatico").
		IssuedAt(time.Now()).
		Subject("juackomdz").
		Claim("user", user).
		Claim("email", email).
		Expiration(time.Now().Add(1 * time.Hour)).
		Build()

	if err != nil {
		log.Print(err)
	}

	signed, errort := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte("secret")), jwt.WithBase64Encoder(base64.URLEncoding))
	if errort != nil {
		log.Print(errort)
	}

	return string(signed)
}

func MiddleJWT() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			h := c.Request().Header.Get("Authorization")
			if len(h) == 0 {
				return c.JSON(401, echo.Map{"mensaje": "Falta token de autorizacion"})
			}

			strToken := strings.TrimSpace(strings.Replace(h, "Bearer", "", 1))

			_, errp := jwt.Parse([]byte(strToken), jwt.WithKey(jwa.HS256(), []byte("secret")), jwt.WithValidate(true), jwt.WithBase64Encoder(base64.URLEncoding))
			if errp != nil {
				log.Print(errp)
				return c.JSON(401, echo.Map{"mensaje": "Error con token, generelo nuevamente"})
			}

			return next(c)
		}
	}
}
