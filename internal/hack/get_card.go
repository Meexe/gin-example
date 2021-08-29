package hack

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (s *Service) GetCard(c *gin.Context) {
	foo := c.Params.ByName("ID")
	ID, err := strconv.Atoi(foo)
	if err != nil {
		c.String(http.StatusBadRequest, "wrong ID format")
		return
	}

	var resp interface{}
	if ID <= 100 {
		resp, err = s.getDepartment(ID)
	} else {
		resp, err = s.getWorker(ID)
	}
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"card": resp})
}

func (s *Service) getDepartment(ID int) (*Department, error) {
	const query = `
		select
			d.id,
			d.department,
			w.id,
			w.name,
			w.position,
			w.phone,
			w.email
		from departments d
		join workers w on d.id = w.department_id
		where d.id = $1
	`

	var (
		dep        Department
		supervisor Worker
	)

	err := s.db.QueryRow(query, ID).Scan(
		&dep.ID,
		&dep.Name,
		&supervisor.ID,
		&supervisor.Name,
		&supervisor.Position,
		&supervisor.Phone,
		&supervisor.Email,
	)
	if err != nil {
		return nil, err
	}

	dep.Supervisor = &supervisor
	dep.Load = rand.Int31n(101)
	dep.Workers, err = s.getWorkers(dep.ID)
	return &dep, err
}

func (s *Service) getWorkers(dep int32) (workers []*Worker, err error) {
	const query = `
		select
			id,
			name,
			position,
			phone,
			email
		from workers
		where not is_supervisor and department_id = $1
	`

	rows, err := s.db.Query(query, dep)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var worker Worker
		worker.Load = rand.Int31n(101)
		if err = rows.Scan(
			&worker.ID,
			&worker.Name,
			&worker.Position,
			&worker.Phone,
			&worker.Email,
		); err != nil {
			return
		}

		worker.Load = rand.Int31n(101)
		workers = append(workers, &worker)
	}
	return
}

func (s *Service) getWorker(ID int) (interface{}, error) {
	const query = `
		select
			w.id,
			w.name,
			w.position,
			w.phone,
			w.email,
			d.department
		from departments d
		join workers w on d.id = w.department_id
		where w.id = $1
	`

	var resp struct {
		ID         int32   `json:"ID"`
		Name       string  `json:"name"`
		Position   string  `json:"position"`
		Phone      string  `json:"phone"`
		Email      string  `json:"email"`
		Department string  `json:"department"`
		Load       int32   `json:"load"`
		Tasks      []*Task `json:"tasks"`
	}

	err := s.db.QueryRow(query, ID).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Position,
		&resp.Phone,
		&resp.Email,
		&resp.Department,
	)
	if err != nil {
		return nil, err
	}

	resp.Load = rand.Int31n(101)
	resp.Tasks, err = s.getWorkerTasks(resp.ID)
	return &resp, err
}

func (s *Service) getWorkerTasks(ID int32) (tasks []*Task, err error) {
	const query = `
		select
			id,
			text,
			status,
			is_regular,
			deadline,
			priority
		from tasks
		where worker_id = $1;
	`

	rows, err := s.db.Query(query, ID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			task Task
			date pq.NullTime
		)
		if err = rows.Scan(
			&task.ID,
			&task.Text,
			&task.Status,
			&task.IsRegular,
			&date,
			&task.Priority,
		); err != nil {
			return
		}
		task.Deadline = date.Time.Format(time.RFC3339)
		tasks = append(tasks, &task)
	}
	return
}
