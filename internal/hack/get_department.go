package hack

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetDepartment(c *gin.Context) {
	foo := c.Params.ByName("ID")
	ID, err := strconv.Atoi(foo)
	if err != nil {
		c.String(http.StatusBadRequest, "wrong ID format")
		return
	}

	dep, err := s.getDepartment(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"department": dep})
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
