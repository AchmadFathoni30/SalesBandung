package controller

import (
	"SalesBandung/config"
	"SalesBandung/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSalesHDR(c *gin.Context) {
	branch_id := c.Query("branch_id")
	trx_date := c.Query("trx_date")

	if branch_id == "" || trx_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID dan Trx Date wajib diisi!"})
		return
	}

	var results []model.SalesHDR
	if err := config.DB.Where("branch_id = ? AND trx_date = ?", branch_id, trx_date).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func GetSalesPAYMENT(c *gin.Context) {
	branch_id := c.Query("branch_id")
	trx_date := c.Query("trx_date")

	if branch_id == "" || trx_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID dan Trx Date wajib diisi!!!"})
		return
	}

	var results []model.SalesPAYMENT
	if err := config.DB.Where("branch_id = ? AND trx_date = ?", branch_id, trx_date).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func GetSalesVOID(c *gin.Context) {
	branch_id := c.Query("branch_id")
	trx_date := c.Query("trx_date")

	if branch_id == "" || trx_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID dan Trx Date harap diisi!!!"})
		return
	}

	var results []model.SalesVOID
	if err := config.DB.Where("branch_id = ? AND trx_date = ?", branch_id, trx_date).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, results)
}
