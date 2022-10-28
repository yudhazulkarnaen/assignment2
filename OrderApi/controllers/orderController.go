package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"assignment2.id/orderapi/database"
	"assignment2.id/orderapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrNotFound error = errors.New("Id tidak ditemukan.")

// DeleteOrder godoc
// @Summary      Delete an order
// @Description  delete order by ID including its items.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderID path uint true "ID number of the order to be deleted."
// @Success      200  {object}  SuccessH
// @Failure      400  {object}  ErrorH
// @Failure      404  {object}  ErrorH
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [delete]
func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := database.DeleteOrderById(uint(parsedID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("id = %d tidak ditemukan.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("id %d terhapus.", parsedID),
	})
}

// UpdateOrder godoc
// @Summary      Update an order
// @Description  update order by ID including its items. Previous items are discarded.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderID path uint true "ID number of the order to be updated."
// @Param        order body models.OrderBody true "JSON of the order to be updated."
// @Success      200  {object}  SuccessH
// @Failure      400  {object}  ErrorH
// @Failure      404  {object}  ErrorH
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [put]
func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var updatedOrder models.Order
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := database.UpdateOrderById(uint(parsedID), &updatedOrder); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("id %d tidak ditemukan.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("id %d terupdate.", parsedID),
	})
}

// CreateOrder godoc
// @Summary      Create an order
// @Description  Create an order including its items, if provided.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order body models.OrderBody true "JSON of the order to be made."
// @Success      201  {object}  models.Order
// @Failure      400  {object}  ErrorH
// @Failure      500  {object}  nil
// @Router       /orders [post]
func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := database.CreateOrder(&newOrder)
	if err != nil {
		if errors.Is(err, models.ErrItemCodeEmpty) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		if errors.Is(err, models.ErrCustomerNameEmpty) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

// GetOrder godoc
// @Summary      Get an order
// @Description  get order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderID path uint true "ID number of the order"
// @Success      200  {object}  models.Order
// @Failure      400  {object}  nil
// @Failure      404  {object}  ErrorH
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [get]
func GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	var orderData models.Order
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	orderData, err = database.GetOrderById(uint(parsedID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("id %d tidak ditemukan.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order": orderData,
	})
}

type ErrorH struct {
	ErrorMessage string `json:"error_message" example:"The error is explained here."`
}
type SuccessH struct {
	Message string `example:"Operation successfull."`
}
