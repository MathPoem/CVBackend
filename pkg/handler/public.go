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
	paramInt := 0
	queryParam := c.Query("universityId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	schools, err := h.service.GetSchool(paramInt)

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
	paramInt := 0
	queryParam := c.Query("schoolId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	departments, err := h.service.GetDepartment(paramInt)

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

func (h *Handler) GetPerson(c *gin.Context) {
	paramInt := 0
	queryParam := c.Query("departmentId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	persons, err := h.service.GetPerson(paramInt)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, persons)

}

func (h *Handler) GetPersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("personId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	person, err := h.service.GetPersonById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *Handler) GetProgram(c *gin.Context) {
	paramInt := 0
	queryParam := c.Query("schoolId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	programs, err := h.service.GetProgram(paramInt)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, programs)

}

func (h *Handler) GetProgramById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("programId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	program, err := h.service.GetProgramById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, program)
}

func (h *Handler) GetCourse(c *gin.Context) {
	paramInt := 0
	queryParam := c.Query("programId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	courses, err := h.service.GetCourse(paramInt)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, courses)

}

func (h *Handler) GetCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	course, err := h.service.GetCourseById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) GetLecture(c *gin.Context) {
	paramInt := 0
	queryParam := c.Query("courseId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	courses, err := h.service.GetLecture(paramInt)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, courses)

}

func (h *Handler) GetLectureById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lectureId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	lecture, err := h.service.GetLectureById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lecture)
}

func (h *Handler) GetSeminar(c *gin.Context) {
	paramInt := 0
	queryParam := c.Query("courseId")

	if len(queryParam) > 0 {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, "invalid university_id param")
			return
		}
		paramInt = id
	}

	seminars, err := h.service.GetSeminar(paramInt)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, seminars)

}

func (h *Handler) GetSeminarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("seminarId"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	seminar, err := h.service.GetSeminarById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, seminar)
}
