package repository

import (
	"CVBackend/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PrivatePostgres struct {
	db *sqlx.DB
}

func NewPrivatePostgres(db *sqlx.DB) *PrivatePostgres {
	return &PrivatePostgres{db: db}
}

func (r *PrivatePostgres) CreateEstimate(estimate models.Estimate) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
		(date_create, user_id, lecture_id, seminar_id, question_1, question_2, question_3, question_4, question_5, question_6, title, description) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`, estimateTable)
	row := r.db.QueryRow(
		query, estimate.Date, estimate.UserID, estimate.LectureID, estimate.SeminarID,
		estimate.Question1, estimate.Question2, estimate.Question3, estimate.Question4, estimate.Question5, estimate.Question6,
		estimate.Title, estimate.Description)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PrivatePostgres) GetEstimate(userId int) ([]models.Estimate, error) {
	var estimates []models.Estimate
	query := fmt.Sprintf(
		"SELECT lecture_id, seminar_id, question_1, question_2, question_3, question_4, question_5, question_6, title, description FROM %s WHERE user_id=$1",
		estimateTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var estimate models.Estimate
		if err := rows.Scan(
			&estimate.LectureID, &estimate.SeminarID, &estimate.Question1, &estimate.Question2, &estimate.Question3,
			&estimate.Question4, &estimate.Question5, &estimate.Question6, &estimate.Title, &estimate.Description,
		); err != nil {
			return nil, err
		}
		estimates = append(estimates, estimate)
	}

	return estimates, nil
}

func (r *PrivatePostgres) CreatePerson(person models.Person, userId int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (department_id, first_name, middle_name, second_name, useradd_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", personTable)
	row := r.db.QueryRow(query, person.DepartmentID, person.FirstName, person.MiddleName, person.SecondName, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
