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

func (r *PublicPostgres) GetDepartmentById(id int) (models.Department, error) {
	var department models.Department
	query := fmt.Sprintf("SELECT id, school_id, name, url FROM %s WHERE id=$1", departmentTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&department.ID, &department.SchoolID, &department.Name, &department.URL); err != nil {
		return models.Department{}, err
	}
	return department, nil
}
