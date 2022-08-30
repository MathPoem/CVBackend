package models

type Estimate struct {
	ID           int    `json:"id" db:"id"`
	Date         string `json:"date_create" db:"date_create"`
	UserID       int    `json:"user_id" db:"user_id"`
	LectureID    int    `json:"lecture_id" db:"lecture_id" binding:"required"`
	SeminarID    int    `json:"seminar_id" db:"seminar_id" binding:"required"`
	QuestionLec1 int    `json:"question_lec_1" db:"question_lec_1" binding:"required"`
	QuestionLec2 int    `json:"question_lec_2" db:"question_lec_2" binding:"required"`
	QuestionLec3 int    `json:"question_lec_3" db:"question_lec_3" binding:"required"`
	QuestionLec4 int    `json:"question_lec_4" db:"question_lec_4" binding:"required"`
	QuestionLec5 int    `json:"question_lec_5" db:"question_lec_5" binding:"required"`
	QuestionSem1 int    `json:"question_sem_1" db:"question_sem_1" binding:"required"`
	QuestionSem2 int    `json:"question_sem_2" db:"question_sem_2" binding:"required"`
	QuestionSem3 int    `json:"question_sem_3" db:"question_sem_3" binding:"required"`
	QuestionSem4 int    `json:"question_sem_4" db:"question_sem_4" binding:"required"`
	QuestionSem5 int    `json:"question_sem_5" db:"question_sem_5" binding:"required"`
	CourseID     int    `json:"course_id" db:"course_id" binding:"required"`
	Title        string `json:"title,omitempty" db:"title"`
	Description  string `json:"description,omitempty" db:"description"`
}

type EstimateResponse struct {
	LectureID    int `json:"lecture_id" db:"lecture_id" binding:"required"`
	SeminarID    int `json:"seminar_id" db:"seminar_id" binding:"required"`
	QuestionLec1 int `json:"question_lec_1" db:"question_lec_1" binding:"required"`
	QuestionLec2 int `json:"question_lec_2" db:"question_lec_2" binding:"required"`
	QuestionLec3 int `json:"question_lec_3" db:"question_lec_3" binding:"required"`
	QuestionLec4 int `json:"question_lec_4" db:"question_lec_4" binding:"required"`
	QuestionLec5 int `json:"question_lec_5" db:"question_lec_5" binding:"required"`
	QuestionSem1 int `json:"question_sem_1" db:"question_sem_1" binding:"required"`
	QuestionSem2 int `json:"question_sem_2" db:"question_sem_2" binding:"required"`
	QuestionSem3 int `json:"question_sem_3" db:"question_sem_3" binding:"required"`
	QuestionSem4 int `json:"question_sem_4" db:"question_sem_4" binding:"required"`
	QuestionSem5 int `json:"question_sem_5" db:"question_sem_5" binding:"required"`
	CourseID     int `json:"course_id" db:"course_id" binding:"required"`
}
