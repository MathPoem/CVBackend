package models

type University struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

type School struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	UniversityID int    `json:"universityID"`
	URL          string `json:"url"`
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SchoolID int    `json:"schoolID"`
	URL      string `json:"url"`
}

type Person struct {
	ID           int    `json:"id"`
	DepartmentID int    `json:"departmentID" binding:"required"`
	FirstName    string `json:"firstName" binding:"required"`
	MiddleName   string `json:"middleName" binding:"required"`
	SecondName   string `json:"secondName" binding:"required"`
	Age          int    `json:"age"`
	URL          string `json:"url"`
	FirstDegree  string `json:"firstDegree"`
	SecondDegree string `json:"secondDegree"`
	ThirdDegree  string `json:"thirdDegree"`
}

type Program struct {
	ID        int    `json:"id"`
	SchoolID  int    `json:"schoolID"`
	Name      string `json:"name"`
	YearStart int    `json:"yearStart"`
	Semester  int    `json:"semester"`
	Url       string `json:"url"`
}

type Course struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Semester            int    `json:"semester"`
	ProgramID           int    `json:"programID"`
	Credits             int    `json:"credits"`
	HoursLecture        int    `json:"hoursLecture"`
	HoursSeminar        int    `json:"hoursSeminar"`
	EstimationInDiploma bool   `json:"estimationInDiploma"`
	Exam                bool   `json:"exam"`
	Test                bool   `json:"test"`
	Description         string `json:"description"`
	URL                 string `json:"url"`
}

type Lecture struct {
	ID       int    `json:"id"`
	Year     int    `json:"year"`
	PersonID int    `json:"pearsonID"`
	CourseID int    `json:"courseID"`
	URL      string `json:"url"`
}

type Seminar struct {
	ID       int    `json:"id"`
	Year     int    `json:"year"`
	PersonID int    `json:"pearsonID"`
	CourseID int    `json:"courseID"`
	URL      string `json:"url"`
}
