package handler

import (
	"golang-gorm-gin/internal/models"
	"golang-gorm-gin/internal/usecase"
	"golang-gorm-gin/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
		return
	}
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) GetLayanan(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	products, err := h.layananUseCase.GetByID(id)
	if err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
			c.AbortWithStatusJSON(http.StatusNotFound, resErr)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : "error occured"})
			return
		}
	}
	response := pkg.NewResponse(http.StatusOK, "Laundry services retrieved successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) CreateLayanan(c *gin.Context){
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.layananUseCase.CreateLayanan(&products) ; err != nil{
		resErr := pkg.NewErrResponse(http.StatusBadRequest, "Can't create service")
		c.AbortWithStatusJSON(http.StatusBadRequest, resErr)
		return
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"sub": products.Id,
	// 	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	// })
	// tokenString, _ := token.SignedString([]byte(env.SECRET))
	
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authoratization", tokenString, 3600 * 24 * 30, "", "", false, true)

	response := pkg.NewResponse(http.StatusCreated, "Laundry services created successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) UpdateLayanan(c *gin.Context){
	var products models.Product
	id, _ := strconv.Atoi(c.Param("id")) 

	if err := c.ShouldBindJSON(&products); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.layananUseCase.UpdateLayanan(id, &products); err!=nil{
		resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
		c.AbortWithStatusJSON(http.StatusNotFound, resErr)
		return
	}
	response := pkg.NewResponse(http.StatusOK, "Laundry services updated successfully", products)
	c.JSON(http.StatusOK, response)
}

func (h *LayananHandler) DeleteLayanan(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.layananUseCase.DeleteLayanan(id); err != nil{
		resErr := pkg.NewErrResponse(http.StatusNotFound, "Laundry service not found")
		c.AbortWithStatusJSON(http.StatusNotFound, resErr)
		return
	}
	response := pkg.NewResponse(http.StatusOK, "Laundry services deleted successfully", nil)
	c.JSON(http.StatusOK, response)
}


