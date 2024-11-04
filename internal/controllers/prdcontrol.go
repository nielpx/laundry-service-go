package controllers

import (
	"golang-gorm-gin/internal/database"
	"golang-gorm-gin/internal/models"
	"golang-gorm-gin/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Index(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	response := pkg.NewIndexResponse(products)
	c.JSON(http.StatusOK, response)
}

func Show(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := database.DB.First(&products, id).Error; err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
			return
		}
	}
	response := pkg.NewShowResponse(products)
	c.JSON(http.StatusOK, response)
}


func Create(c *gin.Context) {
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	database.DB.Create(&products)
	response := pkg.NewCreateResponse(products)
	c.JSON(http.StatusOK, response)

}

func Update(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if database.DB.Model(&products).Where("id = ?", id).Updates(&products).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "not able to change"})
		return
	}
	response := pkg.NewUpdateResponse(products)
	c.JSON(http.StatusOK, response)
}

func Delete(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := database.DB.First(&product, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        } else {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan"})
        }
        return
    }
    if err := database.DB.Delete(&product).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
        return
    }

	response := pkg.NewDeleteResponse(product)
	c.JSON(http.StatusOK, response)
}

 
