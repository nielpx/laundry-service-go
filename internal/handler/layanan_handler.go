package handler

import (
	"golang-gorm-gin/internal/models"
	"golang-gorm-gin/internal/usecase"
	"golang-gorm-gin/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LayananHandler struct {
	layananUseCase usecase.LayananUsecase
}

func NewLayananHandler(usecase usecase.LayananUsecase) *LayananHandler{
	return &LayananHandler{usecase}
}

func (h *LayananHandler) ListLayanan(c *gin.Context){
	products, err := h.layananUseCase.GetAll()
	if err != nil{
		logrus.Error("Error occured")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
		return
	}
	logrus.Info("Laundry services retrieved successfully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) GetLayanan(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	products, err := h.layananUseCase.GetByID(id)
	if err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.Error("Not Found")
			resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
			c.AbortWithStatusJSON(http.StatusNotFound, resErr)
			return
		default:
			logrus.Error("Error occured")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
			return
		}
	}
	logrus.Info("Laundry services retrieved successfully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) CreateLayanan(c *gin.Context){
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil{
		logrus.Error("Something went wrong")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.layananUseCase.CreateLayanan(&products) ; err != nil{
		logrus.Error("Can't create service")
		resErr := pkg.NewErrResponse(http.StatusBadRequest, "Can't create service")
		c.AbortWithStatusJSON(http.StatusBadRequest, resErr)
		return
	}
	logrus.Info("Laundry service created succesfully")
	response := pkg.NewResponse(http.StatusCreated, "Laundry services created successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) UpdateLayanan(c *gin.Context){
	var products models.Product
	id, _ := strconv.Atoi(c.Param("id")) 

	if err := c.ShouldBindJSON(&products); err != nil{
		logrus.Error("Something went wrong")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.layananUseCase.UpdateLayanan(id, &products); err!=nil{
		logrus.Error("Not found")
		resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
		c.AbortWithStatusJSON(http.StatusNotFound, resErr)
		return
	}
	logrus.Info("Laundry services updated succesfully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services updated successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) DeleteLayanan(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.layananUseCase.DeleteLayanan(id); err != nil{
		logrus.Error("Not found")
		resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
		c.AbortWithStatusJSON(http.StatusNotFound, resErr)
		return
	}
	logrus.Info("laundry services deleted succesfuully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services deleted successfully", nil)
	c.JSON(http.StatusOK, response)
}




