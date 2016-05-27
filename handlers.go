package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome!")
// }

func AlgoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	algos, err := RepoFindAll()
	if err != nil {
		log.Println(err)
	}
	log.Println("algoIndex: ", algos)
	if err := json.NewEncoder(w).Encode(algos); err != nil {
		log.Println(err)
	}
}

func AlgoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	algoId := vars["algoId"]
	id, err := strconv.Atoi(algoId)
	if err != nil {
		log.Println(err)
	}
	a, err := RepoFindAlgo(id)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusFound)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Println(err)
	}
}

func AlgoCreate(w http.ResponseWriter, r *http.Request) {
	var algo Algo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &algo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	a, err := RepoCreateAlgo(algo)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func AlgoUpdate(w http.ResponseWriter, r *http.Request) {
	var algo Algo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &algo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	a, err := RepoUpdateAlgo(algo)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func AlgoRandom(w http.ResponseWriter, r *http.Request) {
	algos, err := RepoFindAll()
	if err != nil {
		panic(err)
	}

	if len(algos) <= 0 {
		log.Printf("AlgoRandom: No algos found, return %v\n", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return
	}

	ri := rand.Intn(len(algos))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(algos[ri]); err != nil {
		panic(err)
	}
}
