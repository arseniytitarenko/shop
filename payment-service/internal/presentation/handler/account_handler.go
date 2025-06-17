package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/internal/application/dto"
	"payment/internal/application/errs"
	"payment/internal/application/port/in"
	"payment/internal/presentation/response"
)

type AccountHandler struct {
	accountUseCase in.AccountUseCase
}

func NewAccountHandler(accountUseCase in.AccountUseCase) *AccountHandler {
	return &AccountHandler{accountUseCase: accountUseCase}
}

func (h *AccountHandler) NewAccount(c *gin.Context) {
	var accountRequest dto.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	err := h.accountUseCase.NewAccount(accountRequest.UserID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	var accountRequest dto.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	account, err := h.accountUseCase.GetAccount(accountRequest.UserID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	accountResponse := dto.AccountResponse{
		UserID:  account.UserID,
		Balance: account.Balance,
	}
	c.JSON(http.StatusOK, accountResponse)
}

func (h *AccountHandler) ReplenishAccount(c *gin.Context) {
	var replenishAccountRequest dto.ReplenishAccountRequest
	if err := c.ShouldBindJSON(&replenishAccountRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	err := h.accountUseCase.ReplenishAccount(replenishAccountRequest.UserID, replenishAccountRequest.Amount)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
