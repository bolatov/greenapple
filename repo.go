package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var currentId int

var algos Algos

var db *sql.DB

// Give us some seed data
func init() {
	db = initDB("algo.db")
	createTable()
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS algo(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		desc TEXT
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func RepoFindAlgo(id int) (Algo, error) {
	log.Printf("RepoFindAlgo(id=%v)\n", id)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var a Algo
	row := db.QueryRow("SELECT id, name, desc FROM algo WHERE id = ?", id)
	err := row.Scan(&a.Id, &a.Name, &a.Desc)
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		return a, err
	}
	return a, nil
}

func RepoCreateAlgo(a Algo) (Algo, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO algo(name, desc) VALUES(?, ?)", a.Name, a.Desc)
	if err != nil {
		return a, err
	}
	return a, nil
}

// func RepoUpdateAlgo(a Algo) (Algo, error) {

// }

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

// func RepoDestroyAlgo(id int) error {
// 	for i, t := range algos {
// 		if t.Id == id {
// 			algos = append(algos[:i], algos[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Algo with id of %d to delete", id)
// }
