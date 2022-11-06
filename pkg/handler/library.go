package handler

import (
	"github.com/Hymiside/stubent-media-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createClass(c *gin.Context) {
	schoolId, err := getSchoolId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var class models.Class
	if err = c.BindJSON(&class); err != nil {
		responseWithError(c, http.StatusBadRequest, "error to parse json")
		return
	}
	if class.LetClass == "" || class.NumClass == 0 {
		responseWithError(c, http.StatusBadRequest, "invalid request")
		return
	}

	class.SchoolId = schoolId
	classId, err := h.services.CreateClass(class)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, map[string]int{"classId": classId})
}
