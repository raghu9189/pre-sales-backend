package main

import (
	"net/http"
	"pre-sales-backend/config"
	"pre-sales-backend/controllers"
	"pre-sales-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&models.Lead{}, &models.Estimate{}, &models.SupplierMaterial{})
}

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		// Lead Management
		apiGroup.POST("/leads", controllers.CreateLead)
		apiGroup.GET("/leads", controllers.ListAllLeads)
		apiGroup.PUT("/leads/:id", controllers.UpdateLead)
		apiGroup.DELETE("/leads/:id", controllers.DeleteLead)

		// Cost Estimation
		apiGroup.POST("/estimates", controllers.CreateEstimate)
		apiGroup.GET("/estimates", controllers.ListAllEstimates)
		apiGroup.GET("/suppliers/:item", controllers.SupplierMaterial)

		// Quoting System
		// apiGroup.POST("/quotes")
		// apiGroup.GET("/quotes")
		// apiGroup.PUT("/quotes/:id")

		// Client Portal
		// apiGroup.POST("/client-quotes")
		// apiGroup.GET("/client-quotes/:id")

		// Proposal Tracking
		// apiGroup.GET("/proposals")
		// apiGroup.POST("/proposals/:id/follow-up")
	}

	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}
