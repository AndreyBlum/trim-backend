package handler

import (
	"net/http"

	"trim/internal/domain"
	"trim/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {

	data, err := service.ListTransactions()

	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func CreateTransaction(c *gin.Context) {

	var t domain.Transaction

	if err := c.BindJSON(&t); err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	service.AddTransaction(&t)

	c.JSON(http.StatusCreated, t)
}