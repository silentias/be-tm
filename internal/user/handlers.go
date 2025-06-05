package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	controller *Controller
}

func NewHandler(controller *Controller) *Handler {
	return &Handler{controller: controller}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	userGroup.POST("/register/email-code/send", h.RequestEmailCode)
}

// @Summary Get verification email code
// @Description get verification email code
// @Tags users
// @Produce json
// @Param request body RequestEmailCodeInput true "Email for send verification code"
// @Router /users/register/email-code/send [post]
func (h *Handler) RequestEmailCode(ctx *gin.Context) {
	var request RequestEmailCodeInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	appErr := h.controller.HandleEmailCodeRequest(request)
	if appErr != nil {
		ctx.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	ctx.JSON(201, gin.H{"message": "Verification code for successful sending"})
}
