package models

type Estimate struct {
	ID          int    `json:"id" db:"id"`
	Date        string `json:"date_create" db:"date_create"`
	UserID      int    `json:"user_id" db:"user_id"`
	LectureID   int    `json:"lecture_id" db:"lecture_id" binding:"required"`
	SeminarID   int    `json:"seminar_id" db:"seminar_id" binding:"required"`
	Question1   int    `json:"question_1" db:"question_1" binding:"required"`
	Question2   int    `json:"question_2" db:"question_2" binding:"required"`
	Question3   int    `json:"question_3" db:"question_3" binding:"required"`
	Question4   int    `json:"question_4" db:"question_4" binding:"required"`
	Question5   int    `json:"question_5" db:"question_5" binding:"required"`
	Question6   int    `json:"question_6" db:"question_6" binding:"required"`
	Title       string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
}
