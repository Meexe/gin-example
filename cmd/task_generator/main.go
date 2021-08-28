package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Meexe/gin-example/internal/hack"
	"github.com/Meexe/gin-example/tools"
)

func main() {
	db, err := tools.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// deps, err := getDepartments(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, dep := range deps {
	// 	err = genTask(db, dep)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	deps, err := getAllDepartments(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, dep := range deps {
		err = genTask(db, dep)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getDepartments(db *sql.DB) (deps []*hack.Department, err error) {
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
		var dep hack.Department
		if err = rows.Scan(&dep.ID, &dep.Name); err != nil {
			return
		}
		deps = append(deps, &dep)
	}
	return
}

func getAllDepartments(db *sql.DB) (deps []*hack.Department, err error) {
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
		var dep hack.Department
		if err = rows.Scan(&dep.ID, &dep.Name); err != nil {
			return
		}
		deps = append(deps, &dep)
	}
	return
}

func genTask(db *sql.DB, dep *hack.Department) (err error) {
	const query = `
		insert into tasks(text, contacts, department_id)
		values ($1, $2, $3);
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
	// for i := 0; i < 2; i++ {
	fmt.Println(dep.Name)

	fmt.Print("Text-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if text == "skip" {
		// continue
		return
	}

	fmt.Print("Contacts-> ")
	contacts, _ := reader.ReadString('\n')
	contacts = strings.Replace(contacts, "\n", "", -1)

	_, err = tx.Exec(query, text, contacts, dep.ID)
	if err != nil {
		return
	}
	// }
	return
}
