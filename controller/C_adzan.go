package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PsaTyrE/dbe_adzan/helper"
	"github.com/PsaTyrE/dbe_adzan/model"
	"github.com/labstack/echo/v4"
)

var adzan []model.Adzan

func Index(c echo.Context) error {
	if err := model.DB.Preload("Kota").Find(&adzan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error(),
		})
	}
	return helper.ResJson(c, http.StatusOK, adzan)
}

// func ShowById(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
// 	}

// 	if err := model.DB.Preload("Kota").First(&adzan, id).Error; err != nil {
// 		switch err {
// 		case gorm.ErrRecordNotFound:
// 			return helper.ErrRes(c, http.StatusBadRequest, "record tidak ditemukan")
// 		default:
// 			return helper.ErrRes(c, http.StatusNotFound, err.Error())
// 		}
// 	}
// 	return helper.ErrRes(c, http.StatusOK, adzan)
// }

func Create(c echo.Context) error {
	var adzan model.Adzan
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&adzan); err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}

	defer c.Request().Body.Close()

	if err := model.DB.Create(&adzan).Error; err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResJson(c, http.StatusOK, adzan)
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}

	var adzan model.Adzan
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&adzan); err != nil {
		return helper.ErrRes(c, http.StatusInternalServerError, err.Error())
	}

	if model.DB.Where("id=?", id).Updates(&adzan).RowsAffected == 0 {
		return helper.ErrRes(c, http.StatusBadRequest, "tidak dapat update")
	}

	adzan.Id = int64(id)

	return helper.ResJson(c, http.StatusOK, adzan)
}

func Delete(c echo.Context) error {
	Input := map[string]string{
		"id": "",
	}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&Input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error(),
		})
	}
	defer c.Request().Body.Close()

	var adzan model.Adzan
	if model.DB.Delete(&adzan, Input["id"]).RowsAffected == 0 {
		return helper.ErrRes(c, http.StatusBadRequest, "Tidak Dapat menghapus")
	}

	res := map[string]string{
		"message": "Data berhasil di hapus",
	}
	return helper.ResJson(c, http.StatusOK, res)
}

func GetAdzanByKota(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	kotaID := c.Param("id")

	kotaIDInt, err := strconv.Atoi(kotaID)
	if err != nil {
		return helper.ErrRes(c, http.StatusBadRequest, "Invalid Kota Id")
	}
	var adzanResult []model.Adzan
	if err := model.DB.Preload("Kota").Where("id_kota = ?", kotaIDInt).Find(&adzanResult).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": "Failed to fetch data from the database",
		})
	}
	return helper.ResJson(c, http.StatusOK, adzanResult)
}
