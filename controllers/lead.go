package controllers

import (
	"log"
	"net/http"
	"pre-sales-backend/config"
	"pre-sales-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateLead(ctx *gin.Context) {
	var lead models.Lead

	ctx.ShouldBindBodyWithJSON(&lead)

	if ctx.ShouldBind(&lead) == nil {
		log.Fatal("error binding payload")
	}
	// Create a Lead
	result := config.DB.Create(&lead)

	// Handle error
	if result.RowsAffected != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Lead inserted successfully
	ctx.JSON(http.StatusCreated, gin.H{
		"data": &lead,
	})
}

func ListAllLeads(ctx *gin.Context) {
	var leads []models.Lead
	config.DB.Find(&leads).Limit(100)
	ctx.JSON(http.StatusOK, gin.H{
		"data": leads,
	})
}

func UpdateLead(ctx *gin.Context) {
	var lead models.Lead

	id := ctx.Param("id")

	// Find the existing lead
	if err := config.DB.First(&lead, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	// Bind the request body to the existing lead struct
	if err := ctx.ShouldBindJSON(&lead); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Save the updated lead
	if err := config.DB.Save(&lead).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// Send success response
	ctx.JSON(http.StatusOK, gin.H{"data": lead})
}

func DeleteLead(ctx *gin.Context) {
	var lead models.Lead

	id := ctx.Param("id")

	// Find the existing lead
	if err := config.DB.First(&lead, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	// Delete a Lead
	if err := config.DB.Unscoped().Delete(&lead).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lead"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Lead deleted successfully"})
}
