package handler

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var data models.School

	if err := c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, "error to parse json")
		return
	}
	if data.Name == "" || data.PhoneNumber == "" || data.Email == "" || data.Password == "" {
		responseWithError(c, http.StatusBadRequest, "error invalid request")
		return
	}
	userId, err := h.services.CreateSchool(data)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, map[string]string{"userId": userId})
}

func (h *Handler) signIn(c *gin.Context) {
	var data models.School

	if err := c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, "error to parse json")
		return
	}
	if data.Email == "" || data.Password == "" {
		responseWithError(c, http.StatusBadRequest, "error invalid request")
		return
	}
	token, err := h.services.GenerateToken(data.Email, data.Password)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, map[string]string{"token": token})
}
