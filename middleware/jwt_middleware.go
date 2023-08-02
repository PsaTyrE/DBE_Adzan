package middleware

import (
	"net/http"

	"github.com/PsaTyrE/dbe_adzan/config"
	"github.com/PsaTyrE/dbe_adzan/helper"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err == http.ErrNoCookie {
			res := map[string]string{
				"massage": "StatusUnauthorized",
			}
			return helper.ResJson(c, http.StatusUnauthorized, res)
		}
		// mengambil token value
		tokenString := cookie.Value

		Claims := &config.JwtClaims{}

		// pasrsing Token JWT
		token, err := jwt.ParseWithClaims(tokenString, Claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				res := map[string]string{
					"massage": "StatusUnauthorized",
				}
				return helper.ResJson(c, http.StatusUnauthorized, res)
			case jwt.ValidationErrorExpired:
				// token expired
				res := map[string]string{
					"massage": "StatusUnauthorized, token expired",
				}
				return helper.ResJson(c, http.StatusUnauthorized, res)
			default:
				res := map[string]string{
					"massage": "StatusUnauthorized",
				}
				return helper.ResJson(c, http.StatusUnauthorized, res)
			}
		}
		if !token.Valid {
			res := map[string]string{
				"massage": "StatusUnauthorized",
			}
			return helper.ResJson(c, http.StatusUnauthorized, res)
		}
		return next(c)
	}
}
