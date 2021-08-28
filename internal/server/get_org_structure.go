package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Department struct {
	ID             int32         `json:"id"`
	Name           string        `json:"name"`
	SubDepartments []*Department `json:"departments"`
}

// GetOrgStructure returns OrgStructure
func (s *Server) GetOrgStructure(c *gin.Context) {
	deps, err := s.getOrgStructure()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"departments": deps})
}

func (s *Server) getOrgStructure() (deps []*Department, err error) {
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
