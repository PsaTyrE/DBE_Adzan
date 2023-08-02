package helper

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ResJson(c echo.Context, code int, payload interface{}) error {
	res, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Add("Content-Type", "application/json")
	c.Response().WriteHeader(code)
	c.Response().Write(res)
	return err
}

func ErrRes(c echo.Context, code int, massage string) error {
	return ResJson(c, code, map[string]string{
		"massage": massage,
	})
}
