package repository

import (
	"CVBackend/models"
	"github.com/jmoiron/sqlx"
)

type DefaultContent struct {
	db *sqlx.DB
}

func NewDefaultContent(db *sqlx.DB) *DefaultContent {
	return &DefaultContent{db: db}
}

func (r *DefaultContent) InDefaultContent() error {
	for _, univ := range models.UnivList {
		r.db.QueryRow(
			"INSERT INTO university (id, country, city, name, url) VALUES ($1, $2, $3, $4, $5)",
			&univ.ID,
			&univ.Country,
			&univ.City,
			&univ.Name,
			&univ.URL)
	}
	for _, school := range models.SchoolList {
		r.db.QueryRow(
			"INSERT INTO school (id, name, university_id, url) VALUES ($1, $2, $3, $4)",
			&school.ID,
			&school.Name,
			&school.UniversityID,
			&school.URL)
	}
	for _, program := range models.ProgramList {
		r.db.QueryRow(
			"INSERT INTO program (id, school_id, name, year_start, semester) VALUES ($1, $2, $3, $4, $5)",
			&program.ID,
			&program.SchoolID,
			&program.Name,
			&program.YearStart,
			&program.Semester)
	}
	for _, course := range models.CourseList {
		r.db.QueryRow(
			"INSERT INTO course (id, url, name,credits,description,estimation_in_diploma,exam,hours_lecture,hours_seminar,program_id, test, semester) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
			&course.ID,
			&course.URL,
			&course.Name,
			&course.Credits,
			&course.Description,
			&course.EstimationInDiploma,
			&course.Exam,
			&course.HoursLecture,
			&course.HoursSeminar,
			&course.ProgramID,
			&course.Test,
			&course.Semester)
	}
	for _, department := range models.DepartmentList {
		r.db.QueryRow(
			"INSERT INTO department (id, name, school_id, url) VALUES ($1, $2, $3, $4)",
			&department.ID,
			&department.Name,
			&department.SchoolID,
			&department.URL)
	}
	for _, person := range models.PersonList {
		r.db.QueryRow(
			"INSERT INTO person (id, department_id, first_name, middle_name, second_name, age, url, first_degree, second_degree, third_degree)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			&person.ID,
			&person.DepartmentID,
			&person.FirstName,
			&person.MiddleName,
			&person.SecondName,
			&person.Age,
			&person.URL,
			&person.FirstDegree,
			&person.SecondDegree,
			&person.ThirdDegree)
	}
	for _, lecture := range models.LectureList {

		row := r.db.QueryRow(
			"INSERT INTO lecture (id, year, person_id, course_id, url) VALUES ($1, $2, $3, $4, $5)",
			&lecture.ID,
			&lecture.Year,
			&lecture.PersonID,
			&lecture.CourseID,
			&lecture.URL)
		if err := row.Err(); err != nil {
			return err
		}
	}
	for _, seminar := range models.SeminarList {
		r.db.QueryRow(
			"INSERT INTO seminar (id, year, person_id, course_id, url) VALUES ($1, $2, $3, $4, $5)",
			&seminar.ID,
			&seminar.Year,
			&seminar.PersonID,
			&seminar.CourseID,
			&seminar.URL)
	}
	return nil
}
