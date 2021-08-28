package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Meexe/gin-example/internal/hack"
	"github.com/Meexe/gin-example/tools/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	// deps, err := getDepartments(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, dep := range deps {
	// 	err = genWorker(db, dep, 3, false) // тут генерируем обычных рабочих
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	deps, err := getAllDepartments(db)
	if err != nil {
		log.Fatal(err)
	}

	var phone = 89031787600
	for _, dep := range deps {
		err = genWorker(db, dep, phone) // тут генерируем начальников
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getDepartments(db *sql.DB) (deps []*hack.Object, err error) {
	const query = `
		select
			id,
			department
		from departments
		where not (
			id = any(
				select distinct parent_id
				from departments
				where parent_id is not null
			)
		) order by id;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dep hack.Object
		if err = rows.Scan(&dep.ID, &dep.Name); err != nil {
			return
		}
		deps = append(deps, &dep)
	}
	return
}

func getAllDepartments(db *sql.DB) (deps []*hack.Object, err error) {
	const query = `
		select
			id,
			department
		from departments
		order by id;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dep hack.Object
		if err = rows.Scan(&dep.ID, &dep.Name); err != nil {
			return
		}
		deps = append(deps, &dep)
	}
	return
}

func genWorker(db *sql.DB, dep *hack.Object, phone int) (err error) {
	const query = `
		insert into workers(
			department_id,
			name,
			phone,
			email,
			is_supervisor,
			position
		)
		values ($1, $2, $3, $4, $5, $6);
	`

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	// for i := 0; i < count; i++ {
	fmt.Println(dep.Name)

	fmt.Print("Name-> ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	if name == "skip" {
		// continue
		return
	}

	fmt.Print("Email-> ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)

	_, err = tx.Exec(query, dep.ID, name, phone, email, true, "Руководитель")
	if err != nil {
		return
	}
	// }
	return
}
