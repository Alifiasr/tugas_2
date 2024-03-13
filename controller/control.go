package controller

import (
	"net/http"
	"strconv"
	"tugas__2/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *OrderDB {
	return &OrderDB{
		DB: db,
	}
}

func (ord *OrderDB) CreateOrders(ctx *gin.Context) {
	var (
		order  model.Order
		result gin.H
	)

	ctx.ShouldBindJSON(&order)
	ord.DB.Create(&order)
	result = gin.H{
		"result": order,
	}

	ctx.JSON(http.StatusOK, result)

}

func (ord *OrderDB) GetOrders(ctx *gin.Context) {
	var (
		orders []model.Order
		result gin.H
	)

	err := ord.DB.Model(&model.Order{}).Preload("Items").Find(&orders).Error
	_ = err

	result = gin.H{
		"result": orders,
	}

	ctx.JSON(http.StatusOK, result)
}

func (ord *OrderDB) UpdateOrders(ctx *gin.Context) {
	var (
		order  model.Order
		result gin.H
	)

	id, _ := strconv.Atoi(ctx.Param("id"))
	order.OrderId = id
	ctx.BindJSON(&order)

	ord.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	ord.DB.Model(&order).Updates(&order)

	result = gin.H{
		"Success": "data update successfully",
		"Result":  order,
	}
	ctx.JSON(http.StatusOK, result)
}

func (ord *OrderDB) DeleteOrders(ctx *gin.Context) {
	var (
		order  model.Order
		result gin.H
	)

	id := ctx.Param("id")
	err := ord.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = ord.DB.Select("Items").Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted successfully",
		}
	}

	ctx.JSON(http.StatusOK, result)

}
