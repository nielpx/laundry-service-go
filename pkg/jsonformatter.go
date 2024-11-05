// controllers/responses.go

package pkg

import (
    "golang-gorm-gin/internal/models"
    "net/http"
)

// Response struct for Index
type IndexResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products []models.Product `json:"data"`
}

// Response struct for Show
type ShowResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  models.Product  `json:"data"`
}

// Response struct for Create
type CreateResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  models.Product  `json:"data"`
}

// Response struct for Update
type UpdateResponse struct {
	Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  models.Product  `json:"data"`
}

// Response struct for Delete
type DeleteResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  *models.Product `json:"data"`
}

type ErrGetResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products *[]models.Product `json:"data"`
}

type ErrUpdateResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products *models.Product `json:"data"`
}
// Response variables for Index
func NewIndexResponse(products []models.Product) IndexResponse {
    return IndexResponse{
        Code:      http.StatusOK,
		Status: "Success",
		Message: "Laundry service retrieved successfully",
        Products: products,
    }
}

// Response variable for Show
func NewShowResponse(product models.Product) ShowResponse {
    return ShowResponse{
        Code:      http.StatusOK,
		Status: "Success",
		Message: "Laundry service retrieved successfully",
        Products: product,
    }
}

// Response variable for Create
func NewCreateResponse(product models.Product) CreateResponse {
    return CreateResponse{
		Code:      http.StatusCreated,
		Status: "Success",
		Message: "Laundry service retrieved successfully",
        Products: product,
    }
}

// Response variable for Update
func NewUpdateResponse(product models.Product) UpdateResponse {
    return UpdateResponse{
		Code:      http.StatusOK,
		Status: "Success",
		Message: "Laundry service retrieved successfully",
        Products: product,
    }
}

func NewDeleteResponse(product models.Product) DeleteResponse {
    return DeleteResponse{
		Code:      http.StatusOK,
		Status: "Success",
		Message: "Laundry service deleted successfully",
        Products: nil,
    }
}

func NewErrGetResponse(product models.Product) ErrGetResponse {
    return ErrGetResponse{
		Code:      http.StatusNotFound,
		Status: "error",
		Message: "Laundry service not found",
        Products: nil,
    }
}

func NewErrUpdateResponse(product models.Product) ErrUpdateResponse {
    return ErrUpdateResponse{
		Code:      http.StatusNotFound,
		Status: "error",
		Message: "Laundry service not found",
        Products: nil,
    }
}


