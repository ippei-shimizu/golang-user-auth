package http

import (
	"golang-user-auth/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}
	r.POST("/register", handler.RegisterUser)
}

type registerRequest struct {
	Email  string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userUsecase.RegisterUser(req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ユーザー登録が完了しました"})
}
