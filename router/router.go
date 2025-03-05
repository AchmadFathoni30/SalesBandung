package router

import (
	"SalesBandung/config"
	"SalesBandung/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authorized := r.Group("/SalesBandung", config.BasicAuthMiddleware())
	{
		authorized.GET("/Saleshdr", controller.GetSalesHDR)
		authorized.GET("/Salesdtl", controller.GetSalesDTL)
		authorized.GET("/Salespayment", controller.GetSalesPAYMENT)
		authorized.GET("/Salesvoid", controller.GetSalesVOID)
	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	return r
}
