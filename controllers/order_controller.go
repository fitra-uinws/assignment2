package controllers

import (
	"assisment2/database"
	"assisment2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddOrder(ctx *gin.Context) {
	//inputData := map[string]interface{}{}
	order := models.Order{}
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	//fmt.Println(order)
	err = database.DB.Create(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"status": "SUCCESS", "message": "sukses menambahkan order"})

}

func DeleteOrder(ctx *gin.Context) {
	//inputData := map[string]interface{}{}
	orderId := ctx.Param("orderId")

	if orderId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": "mohon masukan order id"})
		return
	}

	if _, err := strconv.Atoi(orderId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": "mohon masukan order id dalam angka"})
		return
	}

	err := database.DB.Exec("DELETE FROM items where order_id =?", orderId).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}
	err = database.DB.Exec("DELETE FROM orders where order_id =?", orderId).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"status": "SUCCESS", "message": fmt.Sprintf("sukses delete order with id %s", orderId)})

}

func EditOrder(ctx *gin.Context) {
	//inputData := map[string]interface{}{}
	orderIdParam := ctx.Param("orderId")

	if orderIdParam == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": "mohon masukan order id"})
		return
	}

	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": "mohon masukan order id dalam angka"})
		return
	}

	order := models.Order{}
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	order.OrderID = uint(orderId)

	for _, item := range order.Items {
		itemInDb := models.Item{}
		err = database.DB.First(&itemInDb, item.ItemID).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
			return
		}

		if itemInDb.OrderID != order.OrderID {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": "Item tidak dalam Order bersangkutan"})
			return
		}

		err = database.DB.Model(&item).Updates(models.Item{ItemCode: item.ItemCode, Description: item.Description, Quantity: item.Quantity}).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
			return
		}
	}
	//fmt.Println(order)
	err = database.DB.Save(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"status": "SUCCESS", "message": "sukses edit order"})

}

func GetOrders(ctx *gin.Context) {
	//inputData := map[string]interface{}{}
	order := []models.Order{}

	//fmt.Println(order)
	err := database.DB.Preload("Items").Find(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{"data": &order})

}
