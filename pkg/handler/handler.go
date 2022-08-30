package handler

import (
	"CVBackend/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin", "authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		ExposeHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin", "authorization"},
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://localhost:3000" {
				return true
			}
			return false
		},
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/refresh", h.Refresh)
		auth.POST("/logout", h.Logout)

	}

	private := router.Group("/user", h.UserIdentity)
	{
		private.POST("/estimate", h.Estimate)
		private.GET("/estimate", h.GetEstimate)
		private.POST("/person", h.CreatePerson)
		auth.POST("/activate", h.Activate)
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

		person := public.Group("/person")
		{
			person.GET("/", h.GetPerson)
			person.GET("/:personId", h.GetPersonById)
		}

		program := public.Group("/program")
		{
			program.GET("/", h.GetProgram)
			program.GET("/:programId", h.GetProgramById)
		}

		course := public.Group("/course")
		{
			course.GET("/", h.GetCourse)
			course.GET("/:courseId", h.GetCourseById)
		}

		lecture := public.Group("/lecture")
		{
			lecture.GET("/", h.GetLecture)
			lecture.GET("/:lectureId", h.GetLectureById)
		}

		seminar := public.Group("/seminar")
		{
			seminar.GET("/", h.GetSeminar)
			seminar.GET("/:seminarId", h.GetSeminarById)
		}
	}

	return router
}
