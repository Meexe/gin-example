package hack

// Object - объект
type Object struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Load        int32     `json:"load"`
	Children    []*Object `json:"children"`
}

// Department - отдел/департамент муниципалитета
type Department struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	Load       int32     `json:"load"`
	Supervisor *Worker   `json:"supervisor"`
	Workers    []*Worker `json:"workers"`
}

// Worker - сотрудник
type Worker struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Phone    string  `json:"phone"`
	Email    string  `json:"email"`
	Load     int32   `json:"load"`
	Tasks    []*Task `json:"tasks"`
}

// Task - заявка/задача
type Task struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
