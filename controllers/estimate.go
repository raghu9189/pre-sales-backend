package controllers

import (
	"log"
	"net/http"
	"pre-sales-backend/config"
	"pre-sales-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateEstimate(ctx *gin.Context) {
	var estimate models.Estimate

	ctx.ShouldBindBodyWithJSON(&estimate)

	if ctx.ShouldBind(&estimate) == nil {
		log.Fatal("error binding payload")
	}
	// Create a estimate
	result := config.DB.Create(&estimate)

	// Handle error
	if result.RowsAffected != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// estimate inserted successfully
	ctx.JSON(http.StatusCreated, gin.H{
		"data": &estimate,
	})
}

func ListAllEstimates(ctx *gin.Context) {
	var estimates []models.Estimate
	config.DB.Find(&estimates).Limit(100)
	ctx.JSON(http.StatusOK, gin.H{
		"data": estimates,
	})
}

func SupplierMaterial(ctx *gin.Context) {
	var supplierItem []models.SupplierMaterial

	item := ctx.Param("item")

	result := config.DB.Where("Item LIKE ?", "%"+item+"%").Find(&supplierItem).Limit(50)

	if result.Error != nil {
		log.Fatal("error while fetching")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "not fetched data",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": supplierItem,
	})

}
