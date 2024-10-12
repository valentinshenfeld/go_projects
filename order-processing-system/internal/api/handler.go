package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// We'll add dependencies here later
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	RegisterHandlers(r, h)
}

// Implement the ServerInterface methods

func (h *Handler) GetOrders(c *gin.Context) {
	// TODO: Implement
	c.JSON(http.StatusOK, []Order{})
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement order creation logic
	order := Order{
		Id:          1,
		CustomerId:  req.CustomerId,
		Status:      "pending",
		TotalAmount: req.TotalAmount,
	}

	c.JSON(http.StatusCreated, order)
}

func (h *Handler) GetOrder(c *gin.Context, id int) {
	// TODO: Implement
	c.JSON(http.StatusOK, Order{Id: id})
}

func (h *Handler) UpdateOrder(c *gin.Context, id int) {
	var req UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement order update logic
	order := Order{
		Id:     id,
		Status: *req.Status,
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) DeleteOrder(c *gin.Context, id int) {
	// TODO: Implement
	c.Status(http.StatusNoContent)
}
