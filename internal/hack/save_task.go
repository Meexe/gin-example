package hack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) SaveTask(c *gin.Context) {
	var body SaveTaskRequest
	err := c.BindJSON(&body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = s.saveTask(&body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

func (s *Service) saveTask(task *SaveTaskRequest) (err error) {
	const query = `
		insert into tasks (
			text,
			contacts,
			is_regular,
			deadline,
			priority,
			complexity,
			worker_id
		) values (
			$1, $2, $3, $4, $5, $6, $7
		);
	`

	worker, err := s.DistributeTask(task.DepartmentID)
	if err != nil {
		return
	}

	_, err = s.db.Exec(query,
		task.Text,
		task.Contacts,
		task.IsRegular,
		task.Deadline,
		task.Priority,
		task.Complexity,
		worker,
	)
	return
}
