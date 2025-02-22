package router

import (
	"SalesBandung/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"Bandung": "Rahasia@2025",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authorized := r.Group("/SalesBandung", BasicAuthMiddleware())
	{
		authorized.GET("/Saleshdr", controller.GetSalesHDR)
		authorized.GET("/Salespayment", controller.GetSalesPAYMENT)
	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	return r
}
