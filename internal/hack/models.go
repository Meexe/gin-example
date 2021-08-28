package hack

// Department - отдел/департамент муниципалитета
type Department struct {
	ID             int32         `json:"id"`
	Name           string        `json:"name"`
	SubDepartments []*Department `json:"departments"`
}
