package http

import (
	"employee-service/internal/domain"
	"employee-service/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeHandler struct{ uc *usecase.EmployeeUsecase }

func NewEmployeeHandler(r *gin.Engine, uc *usecase.EmployeeUsecase, auth gin.HandlerFunc) {
	h := &EmployeeHandler{uc: uc}
	g := r.Group("/v1/employees")
	g.Use(auth)
	g.POST("", h.create)
	g.GET("/:id", h.get)
	g.PUT("/:id", h.update)
}

type createRequest struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password" binding:"required,min=6"`
	AccessLevel  string  `json:"access_level" binding:"required,oneof=ordinary supervisor manager hr intern"`
	Position     string  `json:"position" binding:"required"`
	WorkLocation string  `json:"work_location"`
	PersonalID   string  `json:"personal_id"`
	Address      string  `json:"address"`
	ZipCode      string  `json:"zip_code"`
	Province     string  `json:"province"`
	City         string  `json:"city"`
	District     string  `json:"district"`
	PhoneNumber  string  `json:"phone_number"`
	PhotoURL     string  `json:"photo_url"`
	NPWP         string  `json:"npwp"`
	BankName     string  `json:"bank_name"`
	BankAccount  string  `json:"bank_account"`
	Salary       float64 `json:"salary"`
}

func (h *EmployeeHandler) create(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.uc.Create(c.Request.Context(), usecase.CreateEmployeeInput{
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		AccessLevel:  domain.AccessLevel(req.AccessLevel),
		Position:     req.Position,
		WorkLocation: req.WorkLocation,
		PersonalID:   req.PersonalID,
		Address:      req.Address,
		ZipCode:      req.ZipCode,
		Province:     req.Province,
		City:         req.City,
		District:     req.District,
		PhoneNumber:  req.PhoneNumber,
		PhotoURL:     req.PhotoURL,
		NPWP:         req.NPWP,
		BankName:     req.BankName,
		BankAccount:  req.BankAccount,
		Salary:       req.Salary,
		Status:       "active",
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         id.String(),
		"created_at": time.Now().In(time.FixedZone("Asia/Jakarta", 7*3600)),
	})
}

func (h *EmployeeHandler) get(c *gin.Context) {
	id := c.Param("id")
	emp, err := h.uc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, emp)
}
func (h *EmployeeHandler) update(c *gin.Context) {
	id := c.Param("id")
	var req domain.Employee
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	req.ID = uid

	if err := h.uc.Update(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee updated"})
}
