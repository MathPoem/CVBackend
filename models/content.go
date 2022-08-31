package models

var (
	UnivList = []University{
		{
			Name:    "МГУ",
			ID:      1,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		//{
		//	Name:    "ВШЭ",
		//	ID:      2,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
		//{
		//	Name:    "МФТИ",
		//	ID:      3,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
		//{
		//	Name:    "МГТУ",
		//	ID:      4,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
		//{
		//	Name:    "МИСИС",
		//	ID:      5,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
		//{
		//	Name:    "МИФИ",
		//	ID:      6,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
		//{
		//	Name:    "МГИМО",
		//	ID:      7,
		//	City:    "Москва",
		//	Country: "Россия",
		//	URL:     "msu.ru",
		//},
	}

	SchoolList = []School{
		{
			ID:           1,
			Name:         "МехМат",
			UniversityID: 1,
			URL:          "math.msu.ru",
		},
		//{
		//	ID:           2,
		//	Name:         "ВМК",
		//	UniversityID: 1,
		//	URL:          "cs.msy.ru",
		//},
		//{
		//	ID:           3,
		//	Name:         "Физфак",
		//	UniversityID: 1,
		//	URL:          "phys.msu.ru",
		//},
		//{
		//	ID:           4,
		//	Name:         "ФКН",
		//	UniversityID: 2,
		//	URL:          "cs.hse",
		//},
		//{
		//	ID:           5,
		//	Name:         "Матфак",
		//	UniversityID: 2,
		//	URL:          "math.hse",
		//},
		//{
		//	ID:           6,
		//	Name:         "Физфак",
		//	UniversityID: 2,
		//	URL:          "phys.hse",
		//},
		//{
		//	ID:           7,
		//	Name:         "ФПМИ",
		//	UniversityID: 3,
		//	URL:          "fpmi.mipt",
		//},
		//{
		//	ID:           8,
		//	Name:         "ЛФИ",
		//	UniversityID: 3,
		//	URL:          "lfi.mipt",
		//},
		//{
		//	ID:           9,
		//	Name:         "ФРКТ",
		//	UniversityID: 3,
		//	URL:          "frkt.mipt",
		//},
	}
	ProgramList = []Program{
		{
			ID:        1,
			SchoolID:  1,
			Name:      "Математика",
			YearStart: 2021,
			Semester:  12,
		},
		{
			ID:        2,
			SchoolID:  1,
			Name:      "Механика",
			YearStart: 2021,
			Semester:  12,
		},
	}
	PersonList = []Person{
		{
			ID:           1,
			DepartmentID: 1,
			FirstName:    "Александр",
			MiddleName:   "Алексеевич",
			SecondName:   "Флеров",
			Age:          33,
			URL:          "https://istina.msu.ru/profile/AAFlerov/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
		{
			ID:           2,
			DepartmentID: 2,
			FirstName:    "Наталия",
			MiddleName:   "Евгеньевна",
			SecondName:   "Шавгулидзе",
			Age:          40,
			URL:          "https://istina.msu.ru/profile/Natalia_Shavgulidze/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
		{
			ID:           3,
			DepartmentID: 2,
			FirstName:    "Юрий",
			MiddleName:   "Викторович",
			SecondName:   "Садовничий",
			Age:          56,
			URL:          "https://istina.msu.ru/profile/uvs/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
		{
			ID:           4,
			DepartmentID: 3,
			FirstName:    "Елена",
			MiddleName:   "Игоревна",
			SecondName:   "Бунина",
			Age:          0,
			URL:          "www.mysite.ru",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		}, {
			ID:           5,
			DepartmentID: 4,
			FirstName:    "Нина",
			MiddleName:   "Аркадьевна",
			SecondName:   "Подольская",
			Age:          0,
			URL:          "http://mech.math.msu.su/~nap/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		}, {
			ID:           6,
			DepartmentID: 4,
			FirstName:    "Михаил",
			MiddleName:   "Андреевич",
			SecondName:   "Ложников",
			Age:          0,
			URL:          "www.mysite.ru",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		}, {
			ID:           7,
			DepartmentID: 4,
			FirstName:    "Илья",
			MiddleName:   "Сергеевич",
			SecondName:   "Григорьев",
			Age:          0,
			URL:          "www.mysite.ru",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		}, {
			ID:           8,
			DepartmentID: 4,
			FirstName:    "Антон",
			MiddleName:   "Вячеславович",
			SecondName:   "Шокуров",
			Age:          0,
			URL:          "http://машинноезрение.рф/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
	}
	DepartmentList = []Department{
		{
			ID:       1,
			Name:     "Кафедра МатематическогоАнализа",
			SchoolID: 1,
			URL:      "1",
		},
		{
			ID:       2,
			Name:     "Кафедра Общей Геометрии и Топологии",
			SchoolID: 1,
			URL:      "1",
		},
		{
			ID:       3,
			Name:     "Кафедра Высшей Алгебры",
			SchoolID: 1,
			URL:      "1",
		},
		{
			ID:       4,
			Name:     "Кафедра Вычислительной Математики",
			SchoolID: 1,
			URL:      "1",
		}, {
			ID:       3,
			Name:     "",
			SchoolID: 1,
			URL:      "1",
		}, {
			ID:       3,
			Name:     "",
			SchoolID: 1,
			URL:      "1",
		}, {
			ID:       3,
			Name:     "",
			SchoolID: 1,
			URL:      "1",
		},
	}
	CourseList = []Course{
		{
			ID:                  1,
			Name:                "Маттематический Анализ 1",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "Первая часть общего курса математического анализа",
			URL:                 "1",
			Test:                true,
			Semester:            1,
		},
		{
			ID:                  2,
			Name:                "Математичексий Анализ 2",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "Вторая часть общего курса математического анализа",
			URL:                 "1",
			Test:                true,
			Semester:            2,
		},
		{
			ID:                  3,
			Name:                "Аналитическая Геометрия",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "Аналичическая геометрия и введение в линейную алгебру",
			URL:                 "1",
			Test:                true,
			Semester:            1,
		},
		{
			ID:                  4,
			Name:                "Элементы Теории Чисел",
			Semester:            1,
			ProgramID:           1,
			Credits:             6,
			HoursLecture:        3,
			HoursSeminar:        3,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Основы теории чисел",
			URL:                 "1",
		},
		{
			ID:                  5,
			Name:                "Алгебра 1",
			Semester:            1,
			ProgramID:           1,
			Credits:             6,
			HoursLecture:        3,
			HoursSeminar:        3,
			EstimationInDiploma: false,
			Exam:                true,
			Test:                true,
			Description:         "Первая часть курса общей алгебры",
			URL:                 "1",
		},
		{
			ID:                  6,
			Name:                "Линейная Алгебра",
			Semester:            2,
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: true,
			Exam:                true,
			Test:                true,
			Description:         "Классический курс линейной алгебры и геометрии",
			URL:                 "1",
		},
		{
			ID:                  7,
			Name:                "Наглядная Геометрия и Топология",
			Semester:            2,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        2,
			HoursSeminar:        2,
			EstimationInDiploma: true,
			Exam:                true,
			Test:                false,
			Description:         "Введение в топологию и общую геометрию",
			URL:                 "1",
		},
		{
			ID:                  8,
			Name:                "Теория Дискретных Функций",
			Semester:            2,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        2,
			HoursSeminar:        2,
			EstimationInDiploma: true,
			Exam:                true,
			Test:                false,
			Description:         "Основы теории дискретных функция и введение в дискретную математику",
			URL:                 "1",
		},
		{
			ID:                  9,
			Name:                "Программирование",
			Semester:            1,
			ProgramID:           1,
			Credits:             6,
			HoursLecture:        3,
			HoursSeminar:        3,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Первый семестровый курс из блока программирование",
			URL:                 "1",
		},
		{
			ID:                  10,
			Name:                "Программирование",
			Semester:            2,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        0,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Первый семестровый курс из блока программирование",
			URL:                 "1",
		},
		{
			ID:                  11,
			Name:                "Английский язык",
			Semester:            2,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        0,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Английский язык для специалистов математиков",
			URL:                 "1",
		},
		{
			ID:                  12,
			Name:                "Английский язык",
			Semester:            3,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        0,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Английский язык для специалистов математиков",
			URL:                 "1",
		},
		{
			ID:                  13,
			Name:                "Алгебра 2",
			Semester:            3,
			ProgramID:           1,
			Credits:             5,
			HoursLecture:        3,
			HoursSeminar:        2,
			EstimationInDiploma: true,
			Exam:                true,
			Test:                true,
			Description:         "Вторая часть курса общей алгебры",
			URL:                 "1",
		},
		{
			ID:                  14,
			Name:                "Математический Анализ 3",
			Semester:            3,
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Test:                true,
			Description:         "Третья часть общего курса математического анализа",
			URL:                 "1",
		},
		{
			ID:                  15,
			Name:                "Введение в Топологию",
			Semester:            3,
			ProgramID:           1,
			Credits:             3,
			HoursLecture:        3,
			HoursSeminar:        0,
			EstimationInDiploma: false,
			Exam:                true,
			Test:                false,
			Description:         "Введение в общую топологию",
			URL:                 "1",
		},
		{
			ID:                  16,
			Name:                "Дифференциальные уравнения",
			Semester:            3,
			ProgramID:           1,
			Credits:             4,
			HoursLecture:        2,
			HoursSeminar:        2,
			EstimationInDiploma: false,
			Exam:                false,
			Test:                true,
			Description:         "Первая часть курса теории решения дифференциальных уравнений",
			URL:                 "1",
		},
		{
			ID:                  17,
			Name:                "Мат Логика и Алгоритмы",
			Semester:            3,
			ProgramID:           1,
			Credits:             3,
			HoursLecture:        3,
			HoursSeminar:        0,
			EstimationInDiploma: false,
			Exam:                true,
			Test:                false,
			Description:         "Введение в математическую логику и теорию алгоритмов",
			URL:                 "1",
		},
		{
			ID:                  18,
			Name:                "Программирование",
			Semester:            3,
			ProgramID:           1,
			Credits:             5,
			HoursLecture:        0,
			HoursSeminar:        5,
			EstimationInDiploma: true,
			Exam:                true,
			Test:                false,
			Description:         "Третий семестровый курс из блока программирование",
			URL:                 "1",
		},
	}
	LectureList = []Lecture{
		{
			ID:       1,
			Year:     2021,
			PersonID: 3,
			CourseID: 3,
			URL:      "1",
		},
	}
	SeminarList = []Seminar{
		{
			ID:       1,
			Year:     2021,
			PersonID: 2,
			CourseID: 3,
			URL:      "1",
		},
	}
)
