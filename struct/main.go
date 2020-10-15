package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	Name   string
	Salary int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type Movie struct {
	Title  string
	Year   int `json: "released"`
	Actors []string
}

func main() {
	//jsonMarshal()
	jsonUnmarshal()
}

func jsonMarshal() {
	var movies = []Movie{
		{Title: "Yol", Year: 1996, Actors: []string{"Yilmaz Guney"}},
		{Title: "Hababam Sinifi", Year: 1990, Actors: []string{"Rifat Ilgaz"}},
	}

	data, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatalf("json marshalling error %s", err)
	}
	fmt.Printf("%s\n", data)
}

func jsonUnmarshal() {
	data := []byte(`[{"Title":"Casablanca"},{"Title":"Cool Hand Luke"},{"Title":"Bullitt"}]`)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshalling failed %s", err)
	}

	fmt.Printf("%+v", titles)
}

func employee() {
	var e = Employee{Name: "Huseyin Altunbay", Salary: 100}
	fmt.Printf("%v", e)
	updateSalary(&e)
	fmt.Printf("%v", e)
}

func wheel() {
	var w Wheel
	w.X = 1
	w.Y = 2
	w.Radius = 4
	w.Spokes = 2
	fmt.Printf("%+v", w)
}

func updateSalary(e *Employee) {
	e.Salary = e.Salary + 1
}
