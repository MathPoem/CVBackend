package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetUniversity(c *gin.Context) {
	universities, err := h.service.GetUniversity()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, universities)
}

func (h *Handler) GetUniversityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("universityId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	universities, err := h.service.GetUniversityById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, universities)
}

func (h *Handler) GetSchool(c *gin.Context) {
	schools, err := h.service.GetSchool()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, schools)

}

func (h *Handler) GetSchoolById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("schoolId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	schools, err := h.service.GetSchoolById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (h *Handler) GetDepartment(c *gin.Context) {
	departments, err := h.service.GetDepartment()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, departments)

}

func (h *Handler) GetDepartmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("departmentId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	departments, err := h.service.GetDepartmentById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, departments)
}
