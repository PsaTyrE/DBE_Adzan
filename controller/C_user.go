package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PsaTyrE/dbe_adzan/config"
	"github.com/PsaTyrE/dbe_adzan/helper"
	"github.com/PsaTyrE/dbe_adzan/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	var Userinput model.User
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&Userinput); err != nil {
		res := map[string]string{
			"message": err.Error(),
		}
		return helper.ResJson(c, http.StatusBadRequest, res)
	}
	defer c.Request().Body.Close()

	// ambil data berdasarkan username
	var user model.User
	if err := model.DB.Where("user_name= ?", Userinput.UserName).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res := map[string]string{
				"message": "user name or password incorect",
			}
			return helper.ResJson(c, http.StatusUnauthorized, res)
		default:
			res := map[string]string{
				"message": err.Error(),
			}
			return helper.ResJson(c, http.StatusInternalServerError, res)
		}
	}

	// verified password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Userinput.Password)); err != nil {
		res := map[string]string{
			"message": "user name or password incorect",
		}
		return helper.ResJson(c, http.StatusUnauthorized, res)
	}

	// proses pembuatan jwt
	expTime := time.Now().Add(time.Minute * 72)
	claims := config.JwtClaims{
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Adzan-App",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarsaikan algo untuk signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	token, err := tokenAlgo.SignedString(config.JWTKey)
	if err != nil {
		res := map[string]string{
			"message": err.Error(),
		}
		return helper.ResJson(c, http.StatusInternalServerError, res)
	}

	// set token ke cooki
	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	res := map[string]string{
		"message": "login berhasil",
	}
	return helper.ResJson(c, http.StatusOK, res)
}
func Register(c echo.Context) error {
	// mengambil Inputan json dari FE

	var Userinput model.User
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&Userinput); err != nil {
		res := map[string]string{
			"message": err.Error(),
		}
		return helper.ResJson(c, http.StatusBadRequest, res)
	}
	defer c.Request().Body.Close()

	// hash password dengan bcrypt

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Userinput.Password), bcrypt.DefaultCost)
	if err != nil {
		res := map[string]string{
			"message": "Failed to hash password",
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	Userinput.Password = string(hashPassword)
	// insert to Db

	if err := model.DB.Create(&Userinput).Error; err != nil {
		res := map[string]string{
			"message": err.Error(),
		}
		return helper.ResJson(c, http.StatusInternalServerError, res)
	}

	res := map[string]string{
		"message": "Register berhasil",
	}
	return helper.ResJson(c, http.StatusOK, res)
}
func Logout(c echo.Context) error {
	// hapus token di cooki
	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	}
	c.SetCookie(cookie)
	res := map[string]string{
		"message": "Logout berhasil",
	}
	return helper.ResJson(c, http.StatusOK, res)
}
