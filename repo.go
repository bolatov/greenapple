package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var currentId int

var db *sql.DB

// Give us some seed data
func initA(dbName string) {
	log.Printf("Init a connection to '%v' db\n", dbName)
	db = initDB(dbName)
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
	a := Algo{}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return a, err
	}

	row := db.QueryRow("SELECT id, name, desc FROM algo WHERE id = ?", id)
	err := row.Scan(&a.Id, &a.Name, &a.Desc)
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
	_, err := db.Exec("INSERT INTO algo(name, desc) VALUES(?, ?)", a.Name, a.Desc)
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

	old.Name, old.Desc = a.Name, a.Desc
	_, err = db.Exec("UPDATE algo SET name=?, desc=? WHERE id=?", old.Name, old.Desc, old.Id)
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
	rows, err := db.Query("SELECT id, name, desc FROM algo")
	if err != nil {
		log.Println(err)
		return algos, err
	}
	defer rows.Close()
	// TODO use technique by Rob Pike (errType)
	for rows.Next() {
		var a Algo
		if err := rows.Scan(&a.Id, &a.Name, &a.Desc); err != nil {
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
