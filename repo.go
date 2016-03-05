package main

import "fmt"

var currentId int

var algos Algos

// Give us some seed data
func init() {
	RepoCreateAlgo(Algo{Name: "Write presentation"})
	RepoCreateAlgo(Algo{Name: "Host meetup"})
}

func RepoFindAlgo(id int) Algo {
	for _, t := range algos {
		if t.Id == id {
			return t
		}
	}
	// return empty Algo if not found
	return Algo{}
}

func RepoCreateAlgo(t Algo) Algo {
	currentId += 1
	t.Id = currentId
	algos = append(algos, t)
	return t
}

func RepoDestroyAlgo(id int) error {
	for i, t := range algos {
		if t.Id == id {
			algos = append(algos[:i], algos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Algo with id of %d to delete", id)
}
