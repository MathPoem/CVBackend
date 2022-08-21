package handler

import (
	"CVBackend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	private := router.Group("/user", h.UserIdentity)
	{
		private.POST("/estimate", h.Estimate)
		private.GET("/estimate", h.GetEstimate)
		private.POST("/person", h.CreatePerson)
	}

	public := router.Group("/api")
	{
		university := public.Group("/university")
		{
			university.GET("/", h.GetUniversity)
			university.GET("/:universityId", h.GetUniversityById)
		}

		school := public.Group("/school")
		{
			school.GET("/", h.GetSchool)
			school.GET("/:schoolId", h.GetSchoolById)
		}

		department := public.Group("/department")
		{
			department.GET("/", h.GetDepartment)
			department.GET("/:departmentId", h.GetDepartmentById)
		}
	}

	return router
}
