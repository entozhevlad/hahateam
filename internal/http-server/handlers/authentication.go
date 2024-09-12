package handlers

import (
	resp "HahaTeam/internal/lib/api/response/registration"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RequestAuth struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type ResponseAuth struct {
	resp.Response
}

type Authentication interface {
	Login(login string, password string) (string, error)
}

func AuthenticationUser(authenticationService Authentication) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestAuth
		if err := c.BindJSON(&req); err != nil {
			slog.Error("Error parsing request", err)
			c.JSON(http.StatusBadRequest, resp.Error("Invalid request"))
			return
		}
		userId, err := authenticationService.Login(req.Username, req.Password)
		if err != nil {
			slog.Error("Error authenticating user", err)
			c.JSON(http.StatusUnauthorized, resp.Error("Unauthorized"))
			return
		}

		slog.Info("Authenticated user", slog.String("userId", userId))

		c.JSON(http.StatusOK, &ResponseAuth{
			Response: *resp.OK(),
		})
	}

}
