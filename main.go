package main

import (
	"pre-sales-backend/config"
	"pre-sales-backend/controllers"
	"pre-sales-backend/models"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&models.Lead{}, &models.Estimate{}, &models.SupplierMaterial{})
}

func main() {
	router := gin.Default()
	router.POST("/api/leads", controllers.CreateLead)
	router.GET("/api/leads", controllers.ListAllLeads)
	router.PUT("/api/leads/:id", controllers.UpdateLead)
	router.DELETE("/api/leads/:id", controllers.DeleteLead)

	router.POST("/api/estimates", controllers.CreateEstimate)
	router.GET("/api/estimates", controllers.ListAllEstimates)
	router.GET("/api/suppliers/:item", controllers.SupplierMaterial)

	// router.POST("/api/quotes")
	// router.GET("/api/quotes")
	// router.PUT("/api/quotes/:id")
	router.Run(":8081")
}
