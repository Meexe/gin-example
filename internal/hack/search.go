package hack

import (
	"database/sql"
	"net/http"

	"github.com/Meexe/gin-example/tools/db"
	"github.com/gin-gonic/gin"
)

func (s *Service) Search(c *gin.Context) {
	searchQuery := c.Query("searchQuery")
	deps, err := s.search(searchQuery)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"objects": deps})
}

func (s *Service) search(searchQuery string) (objs []*Object, err error) {
	const query = `
		select
			d.id,
			d.department as name,
			w.name as description
		from departments d
		join workers w on d.id = w.department_id
		where 
			w.is_supervisor and
			to_tsvector('simple', d.department) @@ to_tsquery('simple', $1)
		union
		select
			id,
			name,
			position as description
		from workers
		where to_tsvector('simple', name) @@ to_tsquery('simple', $1);
	`

	rows, err := s.db.Query(query, db.ToTSQuery(searchQuery))
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var obj Object
		if err = rows.Scan(&obj.ID, &obj.Name, &obj.Description); err != nil {
			return
		}
		objs = append(objs, &obj)
	}
	return
}
