package hack

func (s *Service) DistributeTask(dep int32) (worker int32, err error) {
	const query = `
		select
			w.id
		from workers w
		join tasks t on t.worker_id = w.id
		where not w.is_supervisor and w.department_id = $1
		group by w.id
		order by sum(t.complexity)
		limit 1;
	`

	err = s.db.QueryRow(query, dep).Scan(&worker)
	return
}
