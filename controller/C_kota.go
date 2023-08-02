package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PsaTyrE/dbe_adzan/helper"
	"github.com/PsaTyrE/dbe_adzan/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var kota []model.Kota

func IndexKota(c echo.Context) error {
	if err := model.DB.Find(&kota).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, kota)
}

func ShowKotaById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var kota model.Kota
	if err := model.DB.First(&kota, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.JSON(http.StatusNotFound, "Data tidak ditemukan")
		default:
			return c.JSON(http.StatusNotFound, err.Error())
		}
	}
	return c.JSON(http.StatusOK, kota)
}

func CreateKota(c echo.Context) error {
	var kota model.Kota

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&kota); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer c.Request().Body.Close()
	if err := model.DB.Create(&kota).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, kota)
}

func UpdateKota(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}

	var kota model.Kota
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&kota); err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}

	if model.DB.Where("id_kota=?", id).Updates(&kota).RowsAffected == 0 {
		return helper.ErrRes(c, http.StatusBadRequest, "tidak dapat update")
	}

	kota.IdKota = int64(id)

	return helper.ResJson(c, http.StatusOK, kota)
}

func DeleteKota(c echo.Context) error {
	Input := map[string]string{
		"id": "",
	}

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&kota); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error(),
		})
	}
	defer c.Request().Body.Close()
	var kota model.Kota
	if model.DB.Delete(&kota, Input["id"]).RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "data tidak dapat dihapus")
	}
	res := map[string]string{
		"massage": "Data berhasil dihapus",
	}
	return c.JSON(http.StatusOK, res)
}
