package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var currentId int

var db *sql.DB

// Give us some seed data
func initA() {
	db = initDB()
	createTable()
}

func initDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS algo(
		id SERIAL,
		name VARCHAR(255) NOT NULL,
		descr TEXT
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func RepoFindAlgo(id int) (Algo, error) {
	log.Printf("RepoFindAlgo(id=%v)\n", id)
	a := Algo{}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return a, err
	}

	row := db.QueryRow("SELECT id, name, descr FROM algo WHERE id = $1", id)
	err := row.Scan(&a.Id, &a.Name, &a.Descr)
	if err != nil {
		if err == sql.ErrNoRows {

		} else {
			log.Println(err)
			return a, err
		}
	}
	return a, nil
}

func RepoCreateAlgo(a Algo) (Algo, error) {
	if err := db.Ping(); err != nil {
		log.Println(err)
		return a, err
	}
	_, err := db.Exec("INSERT INTO algo(name, descr) VALUES($1, $2)", a.Name, a.Descr)
	if err != nil {
		return a, err
	}
	a.Id = -1
	return a, nil
}

func RepoUpdateAlgo(a Algo) (Algo, error) {
	if err := db.Ping(); err != nil {
		log.Println(err)
		return a, err
	}
	old, err := RepoFindAlgo(a.Id)
	if err != nil {
		log.Println(err)
		return a, err
	}

	old.Name, old.Descr = a.Name, a.Descr
	_, err = db.Exec("UPDATE algo SET name=$1, descr=$2 WHERE id=$3", old.Name, old.Descr, old.Id)
	return old, err
}

func RepoAlgoCount() (int, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECT count(id) FROM algo")
	var res = -1
	err := row.Scan(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func RepoFindAll() ([]Algo, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	algos := make([]Algo, 0)
	rows, err := db.Query("SELECT id, name, descr FROM algo")
	if err != nil {
		log.Println(err)
		return algos, err
	}
	defer rows.Close()
	// TODO use technique by Rob Pike (errType)
	for rows.Next() {
		var a Algo
		if err := rows.Scan(&a.Id, &a.Name, &a.Descr); err != nil {
			log.Println(err)
		}
		algos = append(algos, a)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return algos, err
	}
	return algos, nil
}

// func RepoDestroyAlgo(id int) error {
// 	for i, t := range algos {
// 		if t.Id == id {
// 			algos = append(algos[:i], algos[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Algo with id of %d to delete", id)
// }
