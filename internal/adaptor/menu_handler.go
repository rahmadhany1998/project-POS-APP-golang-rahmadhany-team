package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	Usecase usecase.MenuUsecase
}

func NewMenuHandler(u usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{Usecase: u}
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.CreateMenu(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": product,
	})
}

func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	menus, err := h.Usecase.GetAllMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch menus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menus})
}

func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid menu id"})
		return
	}

	uintID := uint(id)

	//validasi untuk ID tidak boleh negatif
	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
	}

	menu, err := h.Usecase.GetMenuByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid menu id"})
		return
	}

	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = id

	if err := h.Usecase.UpdateMenu(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu updated successfully"})
	return
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	uintID := uint(id)

	err = h.Usecase.DeleteMenu(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu deleted successfully"})
}
