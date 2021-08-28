package hack

import (
	"database/sql"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrgStructure возвращает орг структуру муниципалитета
func (s *Service) GetOrgStructure(c *gin.Context) {
	objs, err := s.getOrgStructure()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"objects": objs})
}

func (s *Service) getOrgStructure() (deps []*Object, err error) {
	const query = `
		select
			d.id,
			d.department,
			w.name,
			coalesce(d.parent_id, 0)
		from departments d
		join workers w on d.id = w.department_id
		where w.is_supervisor
		order by d.id;
	`

	var (
		parentID   int32
		objMapping = make(map[int32]*Object)
		workers    []*Object
	)

	rows, err := s.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var obj Object
		obj.Load = rand.Int31n(101)
		if err = rows.Scan(&obj.ID, &obj.Name, &obj.Description, &parentID); err != nil {
			return
		}

		if el, ok := objMapping[obj.ID]; !ok {
			objMapping[obj.ID] = &obj
		} else {
			el.Name = obj.Name
		}

		if parentID == 0 {
			deps = append(deps, &obj)
		} else if el, ok := objMapping[parentID]; ok {
			el.Children = append(el.Children, &obj)
		} else {
			el = &Object{ID: parentID}
			objMapping[parentID] = el
			el.Children = append(el.Children, &obj)
		}

		if workers, err = s.getOrgStructureWorkers(obj.ID); err != nil {
			return
		} else if len(workers) > 0 {
			obj.Children = workers
		}
	}
	return
}

func (s *Service) getOrgStructureWorkers(dep int32) (workers []*Object, err error) {
	const query = `
		select
			id,
			name,
			position
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
		var obj Object
		obj.Load = rand.Int31n(101)
		if err = rows.Scan(&obj.ID, &obj.Name, &obj.Description); err != nil {
			return
		}
		workers = append(workers, &obj)
	}
	return
}
