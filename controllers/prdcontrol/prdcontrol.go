package prdcontrol

import (
	"golang-gorm-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Create(c *gin.Context) {
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&products)
	c.JSON(http.StatusOK, gin.H{"products" : products})

}

func Update(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&products).Where("id = ?", id).Updates(&products).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "not able to change"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : products})
}

func Delete(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&products).Where("id = ?", id).Delete(&products).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "not able to change"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : products})
}

 
