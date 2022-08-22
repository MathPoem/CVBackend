package repository

import (
	"CVBackend/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Default interface {
	InDefaultContent() error
}

type Private interface {
	CreateEstimate(estimate models.Estimate) (int, error)
	GetEstimate(userId int) ([]models.Estimate, error)
	CreatePerson(person models.Person, userId int) (int, error)
}

type Public interface {
	GetUniversity() ([]models.University, error)
	GetUniversityById(id int) (models.University, error)
	GetSchool() ([]models.School, error)
	GetSchoolByUniversity(id int) ([]models.School, error)
	GetSchoolById(id int) (models.School, error)
	GetDepartment() ([]models.Department, error)
	GetDepartmentBySchool(id int) ([]models.Department, error)
	GetDepartmentById(id int) (models.Department, error)
	GetPerson() ([]models.Person, error)
	GetPersonByDepartment(id int) ([]models.Person, error)
	GetPersonById(id int) (models.Person, error)
	GetProgram() ([]models.Program, error)
	GetProgramBySchool(id int) ([]models.Program, error)
	GetProgramById(id int) (models.Program, error)
	GetCourse() ([]models.Course, error)
	GetCourseByProgram(id int) ([]models.Course, error)
	GetCourseById(id int) (models.Course, error)
	GetLecture() ([]models.Lecture, error)
	GetLectureByCourse(id int) ([]models.Lecture, error)
	GetLectureById(id int) (models.Lecture, error)
	GetSeminar() ([]models.Seminar, error)
	GetSeminarByCourse(id int) ([]models.Seminar, error)
	GetSeminarById(id int) (models.Seminar, error)
}

type Repository struct {
	Authorization
	Default
	Private
	Public
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Default:       NewDefaultContent(db),
		Private:       NewPrivatePostgres(db),
		Public:        NewPublicPostgres(db),
	}
}
