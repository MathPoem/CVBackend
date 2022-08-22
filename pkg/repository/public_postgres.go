package repository

import (
	"CVBackend/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PublicPostgres struct {
	db *sqlx.DB
}

func NewPublicPostgres(db *sqlx.DB) *PublicPostgres {
	return &PublicPostgres{db: db}
}

func (r *PublicPostgres) GetUniversity() ([]models.University, error) {
	var universities []models.University
	query := fmt.Sprintf("SELECT id, country, city, name, url FROM %s", universityTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var university models.University
		if err := rows.Scan(
			&university.ID, &university.Country, &university.City, &university.Name, &university.URL); err != nil {
			return nil, err
		}
		universities = append(universities, university)
	}
	return universities, nil
}

func (r *PublicPostgres) GetUniversityById(id int) (models.University, error) {
	var university models.University
	query := fmt.Sprintf("SELECT id, country, city, name, url FROM %s WHERE id=$1", universityTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&university.ID, &university.Country, &university.City, &university.Name, &university.URL); err != nil {
		return models.University{}, err
	}
	return university, nil
}

func (r *PublicPostgres) GetSchool() ([]models.School, error) {
	var schools []models.School
	query := fmt.Sprintf("SELECT id, university_id, name, url FROM %s", schoolTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var school models.School
		if err := rows.Scan(
			&school.ID, &school.UniversityID, &school.Name, &school.URL); err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	return schools, nil
}

func (r *PublicPostgres) GetSchoolByUniversity(id int) ([]models.School, error) {
	var schools []models.School
	query := fmt.Sprintf("SELECT id, university_id, name, url FROM %s WHERE university_id=$1", schoolTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var school models.School
		if err := rows.Scan(
			&school.ID, &school.UniversityID, &school.Name, &school.URL); err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	return schools, nil
}

func (r *PublicPostgres) GetSchoolById(id int) (models.School, error) {
	var school models.School
	query := fmt.Sprintf("SELECT id, university_id, name, url FROM %s WHERE id=$1", schoolTable)
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&school.ID, &school.UniversityID, &school.Name, &school.URL); err != nil {
		return models.School{}, err
	}
	return school, nil
}

func (r *PublicPostgres) GetDepartment() ([]models.Department, error) {
	var departments []models.Department
	query := fmt.Sprintf("SELECT id, school_id, name, url FROM %s", departmentTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var department models.Department
		if err := rows.Scan(
			&department.ID, &department.SchoolID, &department.Name, &department.URL); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}

func (r *PublicPostgres) GetDepartmentBySchool(id int) ([]models.Department, error) {
	var departments []models.Department
	query := fmt.Sprintf("SELECT id, school_id, name, url FROM %s WHERE school_id=$1", departmentTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var department models.Department
		if err := rows.Scan(
			&department.ID, &department.SchoolID, &department.Name, &department.URL); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}

func (r *PublicPostgres) GetDepartmentById(id int) (models.Department, error) {
	var department models.Department
	query := fmt.Sprintf("SELECT id, school_id, name, url FROM %s WHERE id=$1", departmentTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&department.ID, &department.SchoolID, &department.Name, &department.URL); err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (r *PublicPostgres) GetPerson() ([]models.Person, error) {
	var persons []models.Person
	query := fmt.Sprintf("SELECT id, department_id, first_name, middle_name, second_name, url FROM %s", personTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(
			&person.ID, &person.DepartmentID, &person.FirstName, &person.MiddleName, &person.SecondName, &person.URL); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (r *PublicPostgres) GetPersonByDepartment(id int) ([]models.Person, error) {
	var persons []models.Person
	query := fmt.Sprintf("SELECT id, department_id, first_name, middle_name, second_name, url FROM %s WHERE department_id=$1", personTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(
			&person.ID, &person.DepartmentID, &person.FirstName, &person.MiddleName, &person.SecondName, &person.URL); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (r *PublicPostgres) GetPersonById(id int) (models.Person, error) {
	var person models.Person
	query := fmt.Sprintf("SELECT id, department_id, first_name, middle_name, second_name, url FROM %s WHERE id=$1", personTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&person.ID, &person.DepartmentID, &person.FirstName, &person.MiddleName, &person.SecondName, &person.URL); err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (r *PublicPostgres) GetProgram() ([]models.Program, error) {
	var programs []models.Program
	query := fmt.Sprintf("SELECT id, school_id, name, year_start, semester, url FROM %s", programTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var program models.Program
		if err := rows.Scan(
			&program.ID, &program.SchoolID, &program.Name, &program.YearStart, &program.Semester, &program.Url); err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (r *PublicPostgres) GetProgramBySchool(id int) ([]models.Program, error) {
	var programs []models.Program
	query := fmt.Sprintf("SELECT id, school_id, name, year_start, semester, url FROM %s WHERE school_id=$1", programTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var program models.Program
		if err := rows.Scan(
			&program.ID, &program.SchoolID, &program.Name, &program.YearStart, &program.Semester, &program.Url); err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (r *PublicPostgres) GetProgramById(id int) (models.Program, error) {
	var program models.Program
	query := fmt.Sprintf("SELECT id, school_id, name, year_start, semester, url FROM %s WHERE id=$1", programTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&program.ID, &program.SchoolID, &program.Name, &program.YearStart, &program.Semester, &program.Url); err != nil {
		return models.Program{}, err
	}
	return program, nil
}

func (r *PublicPostgres) GetCourse() ([]models.Course, error) {
	var courses []models.Course
	query := fmt.Sprintf("SELECT id, program_id, name, hours_lecture, hours_seminar, estimation_in_diploma, exam, description, url FROM %s", courseTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(
			&course.ID, &course.ProgramID, &course.Name, &course.HoursLecture, &course.HoursSeminar,
			&course.EstimationInDiploma, &course.Exam, &course.Description, &course.URL); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *PublicPostgres) GetCourseByProgram(id int) ([]models.Course, error) {
	var courses []models.Course
	query := fmt.Sprintf("SELECT id, program_id, name, hours_lecture, hours_seminar, estimation_in_diploma, exam, description, url FROM %s WHERE program_id=$1",
		courseTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(
			&course.ID, &course.ProgramID, &course.Name, &course.HoursLecture, &course.HoursSeminar,
			&course.EstimationInDiploma, &course.Exam, &course.Description, &course.URL); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *PublicPostgres) GetCourseById(id int) (models.Course, error) {
	var course models.Course
	query := fmt.Sprintf("SELECT id, program_id, name, hours_lecture, hours_seminar, estimation_in_diploma, exam, description, url FROM %s WHERE id=$1",
		courseTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&course.ID, &course.ProgramID, &course.Name, &course.HoursLecture, &course.HoursSeminar,
		&course.EstimationInDiploma, &course.Exam, &course.Description, &course.URL); err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (r *PublicPostgres) GetLecture() ([]models.Lecture, error) {
	var lectures []models.Lecture
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s", lectureTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var lecture models.Lecture
		if err := rows.Scan(
			&lecture.ID, &lecture.PersonID, &lecture.CourseID, &lecture.Year, &lecture.URL); err != nil {
			return nil, err
		}
		lectures = append(lectures, lecture)
	}
	return lectures, nil
}

func (r *PublicPostgres) GetLectureByCourse(id int) ([]models.Lecture, error) {
	var lectures []models.Lecture
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s WHERE course_id=$1", lectureTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var lecture models.Lecture
		if err := rows.Scan(
			&lecture.ID, &lecture.PersonID, &lecture.CourseID, &lecture.Year, &lecture.URL); err != nil {
			return nil, err
		}
		lectures = append(lectures, lecture)
	}
	return lectures, nil
}

func (r *PublicPostgres) GetLectureById(id int) (models.Lecture, error) {
	var lecture models.Lecture
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s WHERE id=$1", lectureTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&lecture.ID, &lecture.PersonID, &lecture.CourseID, &lecture.Year, &lecture.URL); err != nil {
		return models.Lecture{}, err
	}
	return lecture, nil
}

func (r *PublicPostgres) GetSeminar() ([]models.Seminar, error) {
	var seminars []models.Seminar
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s", seminarTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var seminar models.Seminar
		if err := rows.Scan(
			&seminar.ID, &seminar.PersonID, &seminar.CourseID, &seminar.Year, &seminar.URL); err != nil {
			return nil, err
		}
		seminars = append(seminars, seminar)
	}
	return seminars, nil
}

func (r *PublicPostgres) GetSeminarByCourse(id int) ([]models.Seminar, error) {
	var seminars []models.Seminar
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s WHERE course_id=$1", seminarTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var seminar models.Seminar
		if err := rows.Scan(
			&seminar.ID, &seminar.PersonID, &seminar.CourseID, &seminar.Year, &seminar.URL); err != nil {
			return nil, err
		}
		seminars = append(seminars, seminar)
	}
	return seminars, nil
}

func (r *PublicPostgres) GetSeminarById(id int) (models.Seminar, error) {
	var seminar models.Seminar
	query := fmt.Sprintf("SELECT id, person_id, course_id, year, url FROM %s WHERE id=$1", seminarTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&seminar.ID, &seminar.PersonID, &seminar.CourseID, &seminar.Year, &seminar.URL); err != nil {
		return models.Seminar{}, err
	}
	return seminar, nil
}
