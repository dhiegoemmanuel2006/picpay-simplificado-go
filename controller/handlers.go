package controller

import (
	"errors"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/pkg/erros"
	request2 "github.com/dhiegoemmanuel2006/picpay-simplificado-go/pkg/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *Api) HandleTranfer(c *gin.Context) {
	var req request2.TranferRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Request does not contain a valid tranfer"})
	}
	if req.Value <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Value is too small"})
	}
	payer, err := api.Us.GetUserByID(req.Payer)
	if err != nil {
		if errors.Is(err, erros.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payer not found"})
		}
	}

	payee, err := api.Us.GetUserByID(req.Payee)
	if err != nil {
		if errors.Is(err, erros.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payee not found"})
			return
		}
	}

	_, err = api.Us.ValidTransaction(payer, payee, req.Value)
	if err != nil {
		if errors.Is(err, erros.ErrLojistaNotCanSendMoney) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Lojista not can send money"})
			return
		}
		if errors.Is(err, erros.ErrPayerDontHaveSufficientBalance) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Payer Dont have sufficient balance"})
			return
		}
		if errors.Is(err, erros.ErrTranferUnauthorized) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Transfer Unauthorized"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected Error"})
	}

	api.Us.DoTransaction(payer, payee, req.Value)
	c.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func (api *Api) HandleCreateUser(c *gin.Context) {
	var req request2.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Request does not contain a valid user"})
		return
	}
	uuid, err := api.Us.CreateUser(req)
	if err != nil {
		if errors.Is(err, erros.ErrThisRoleIsNotAllowed) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This role is not allowed to create user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "id": uuid})
}
