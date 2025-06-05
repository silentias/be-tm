package user

import "be-tm/pkg/common"

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

type RequestEmailCodeInput struct {
	Email string `json:"email" binding:"required,email" example:"user@example.com"`
}

func (c *Controller) HandleEmailCodeRequest(request RequestEmailCodeInput) *common.AppError {
	err := c.service.GenerateEmailCode(request.Email)
	if err != nil {
		return common.InternalServerError()
	}
	return nil
}
