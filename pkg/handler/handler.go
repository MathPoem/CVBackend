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

	//public := router.Group("/api")
	//{
	//	university := public.Group("/university")
	//	{
	//		university.GET("/")
	//		university.GET("/:universityId")
	//	}
	//	school := public.Group("/school")
	//	{
	//		school.GET("/")
	//		school.GET("/:schoolId")
	//	}
	//	department := public.Group("/department")
	//	{
	//		department.GET("/")
	//		department.GET("/:departmentId")
	//	}
	//}

	return router
}
