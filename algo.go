package main

type Algo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type Algos []Algo
