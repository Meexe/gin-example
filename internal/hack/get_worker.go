package hack

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetWorker(c *gin.Context) {
	foo := c.Params.ByName("ID")
	ID, err := strconv.Atoi(foo)
	if err != nil {
		c.String(http.StatusBadRequest, "wrong ID format")
		return
	}

	worker, err := s.getWorker(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"worker": worker})
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
		ID         int32
		Name       string
		Position   string
		Phone      string
		Email      string
		Department string
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

	return &resp, err
}
