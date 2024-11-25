package handler

import (
	// "golang-gorm-gin/internal/models"
	"golang-gorm-gin/internal/models"
	"golang-gorm-gin/internal/usecase"
	"golang-gorm-gin/pkg"
	"net/http"
	"strconv"

	_ "golang-gorm-gin/cmd/app/docs"

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

// ListLayanan godoc
// @Summary Get all laundry services
// @Description Retrieve a list of all laundry services
// @Tags LaundryServices
// @Accept  json
// @Produce  json
// @Success 200 {object} pkg.SuccessResponse
// @Failure 500 {object} pkg.NilResponse
// @Router /laundry-services [get]
func (h *LayananHandler) ListLayanan(c *gin.Context){
	products, err := h.layananUseCase.GetAll()
	if err != nil{
		logrus.Error("Error occured")
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message" : "error occured"})
		return
	}
	logrus.Info("Laundry services retrieved successfully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

// GetLayanan godoc
// @Summary Get a laundry service by ID
// @Description Retrieve a specific laundry service by its ID
// @Tags LaundryServices
// @Accept  json
// @Produce  json
// @Param id path int true "Laundry Service ID"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 404 {object} pkg.NilResponse
// @Failure 500 {object} pkg.NilResponse
// @Router /laundry-services/{id} [get]
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message" : "error occured"})
			return
		}
	}
	logrus.Info("Laundry services retrieved successfully")
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

// CreateLayanan godoc
// @Summary Create a new laundry service
// @Description Add a new laundry service
// @Tags LaundryServices
// @Accept  json
// @Produce  json
// @Param request body models.Product true "Laundry Service Data"
// @Success 201 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.NilResponse
// @Failure 500 {object} pkg.NilResponse
// @Router /laundry-services [post]
func (h *LayananHandler) CreateLayanan(c *gin.Context){
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil{
		logrus.Error("Something went wrong")
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
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

// UpdateLayanan godoc
// @Summary Update an existing laundry service
// @Description Update laundry service by ID
// @Tags LaundryServices
// @Accept  json
// @Produce  json
// @Param id path int true "Laundry Service ID"
// @Param request body models.Product true "Updated Laundry Service Data"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.NilResponse
// @Failure 404 {object} pkg.NilResponse
// @Router /laundry-services/{id} [put]
func (h *LayananHandler) UpdateLayanan(c *gin.Context){
	var products models.Product
	id, _ := strconv.Atoi(c.Param("id")) 

	if err := c.ShouldBindJSON(&products); err != nil{
		logrus.Error("Something went wrong")
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
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

// DeleteLayanan godoc
// @Summary Delete a laundry service
// @Description Delete a specific laundry service by ID
// @Tags LaundryServices
// @Accept  json
// @Produce  json
// @Param id path int true "Laundry Service ID"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 404 {object} pkg.NilResponse
// @Router /laundry-services/{id} [delete]
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




