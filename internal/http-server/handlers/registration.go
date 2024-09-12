package handlers

import (
	resp "HahaTeam/internal/lib/api/response/registration"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RequestReg struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Company  string `json:"company"`
}

type Response struct {
	resp.Response
}

type Registration interface {
	Register(login string, password string, company string) (string, error)
}

func CreateNewUser(registrationService Registration) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestReg

		if err := c.BindJSON(&req); err != nil {
			slog.Error("Error parsing request", err)
			c.JSON(http.StatusBadRequest, resp.Error("Invalid request"))
			return
		}

		userId, err := registrationService.Register(req.Login, req.Password, req.Company)
		if err != nil {
			slog.Error("Error registering user", err)
			c.JSON(http.StatusInternalServerError, resp.Error("Failed to register user"))
			return
		}

		slog.Info("User created", slog.String("userId", userId))

		c.JSON(http.StatusOK, &Response{
			Response: *resp.OK(),
		})
	}
}
