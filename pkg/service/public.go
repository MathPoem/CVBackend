package service

import (
	"CVBackend/models"
	"CVBackend/pkg/repository"
)

type PublicService struct {
	repo repository.Public
}

func NewPublicService(repo repository.Public) *PublicService {
	return &PublicService{repo: repo}
}

func (r *PublicService) GetUniversity() ([]models.University, error) {
	return r.repo.GetUniversity()
}

func (r *PublicService) GetUniversityById(id int) (models.University, error) {
	return r.repo.GetUniversityById(id)
}

func (r *PublicService) GetSchool(id int) ([]models.School, error) {
	if id > 0 {
		return r.repo.GetSchoolByUniversity(id)
	} else {
		return r.repo.GetSchool()
	}
}

func (r *PublicService) GetSchoolById(id int) (models.School, error) {
	return r.repo.GetSchoolById(id)
}

func (r *PublicService) GetDepartment(id int) ([]models.Department, error) {
	if id > 0 {
		return r.repo.GetDepartmentBySchool(id)
	} else {
		return r.repo.GetDepartment()
	}
}

func (r *PublicService) GetDepartmentById(id int) (models.Department, error) {
	return r.repo.GetDepartmentById(id)
}

func (r *PublicService) GetPerson(id int) ([]models.Person, error) {
	if id > 0 {
		return r.repo.GetPersonByDepartment(id)
	} else {
		return r.repo.GetPerson()
	}
}

func (r *PublicService) GetPersonById(id int) (models.Person, error) {
	return r.repo.GetPersonById(id)
}

func (r *PublicService) GetProgram(id int) ([]models.Program, error) {
	if id > 0 {
		return r.repo.GetProgramBySchool(id)
	} else {
		return r.repo.GetProgram()
	}
}

func (r *PublicService) GetProgramById(id int) (models.Program, error) {
	return r.repo.GetProgramById(id)
}

func (r *PublicService) GetCourse(id int) ([]models.Course, error) {
	if id > 0 {
		return r.repo.GetCourseByProgram(id)
	} else {
		return r.repo.GetCourse()
	}
}

func (r *PublicService) GetCourseById(id int) (models.Course, error) {
	return r.repo.GetCourseById(id)
}

func (r *PublicService) GetLecture(id int) ([]models.Lecture, error) {
	if id > 0 {
		return r.repo.GetLectureByCourse(id)
	} else {
		return r.repo.GetLecture()
	}
}

func (r *PublicService) GetLectureById(id int) (models.Lecture, error) {
	return r.repo.GetLectureById(id)
}

func (r *PublicService) GetSeminar(id int) ([]models.Seminar, error) {
	if id > 0 {
		return r.repo.GetSeminarByCourse(id)
	} else {
		return r.repo.GetSeminar()
	}
}

func (r *PublicService) GetSeminarById(id int) (models.Seminar, error) {
	return r.repo.GetSeminarById(id)
}
