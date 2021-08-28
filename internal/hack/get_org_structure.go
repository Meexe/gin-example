package hack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrgStructure возвращает орг структуру муниципалитета
func (s *Service) GetOrgStructure(c *gin.Context) {
	deps, err := s.getOrgStructure()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"departments": deps})
}

func (s *Service) getOrgStructure() (deps []*Department, err error) {
	const query = `
		select
			id,
			department,
			coalesce(parent_id, 0)
		from departments
		order by id;
	`

	var (
		parentID   int32
		depMapping = make(map[int32]*Department)
	)

	rows, err := s.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {

		var dep Department
		if err = rows.Scan(&dep.ID, &dep.Name, &parentID); err != nil {
			return
		}

		if el, ok := depMapping[dep.ID]; !ok {
			depMapping[dep.ID] = &dep
		} else {
			el.Name = dep.Name
		}

		if parentID == 0 {
			deps = append(deps, &dep)
		} else if el, ok := depMapping[parentID]; ok {
			el.SubDepartments = append(el.SubDepartments, &dep)
		} else {
			el = &Department{ID: parentID}
			depMapping[parentID] = el
			el.SubDepartments = append(el.SubDepartments, &dep)
		}

	}
	return
}
