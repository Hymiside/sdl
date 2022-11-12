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

func (h *Handler) createStudent(c *gin.Context) {
	schoolId, err := getSchoolId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var student models.Student
	if err = c.BindJSON(&student); err != nil {
		responseWithError(c, http.StatusBadRequest, "error to parse json")
		return
	}
	if student.FirstName == "" || student.LastName == "" || student.MiddleName == "" || student.ClassNum == 0 ||
		student.ClassLet == "" || student.Email == "" || student.PhoneNumber == "" {
		responseWithError(c, http.StatusBadRequest, "invalid request")
		return
	}

	student.SchoolId = schoolId
	studentId, err := h.services.CreateStudent(student)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, map[string]int{"studentId": studentId})
}

func (h *Handler) getAllClasses(c *gin.Context) {
	schoolId, err := getSchoolId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	classes, err := h.services.GetAllClasses(schoolId)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, classes)
}

func (h *Handler) getAllStudents(c *gin.Context) {
	schoolId, err := getSchoolId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	students, err := h.services.GetAllStudents(schoolId)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, students)
}
