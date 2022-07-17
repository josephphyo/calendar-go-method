package main

import (
	"fmt"
	"log"

	"github.com/josephphyo/export-packages-test/setter-methods/calendar"
)

func main() {
	event := calendar.Event{}

	error := event.SetTitle("Some Event!!")
	if error != nil {
		log.Fatal(error)
	}

	error = event.SetYear(2022)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Year:", event.Year())

	error = event.SetMonth(02)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Month:", event.Month())

	error = event.SetDay(14)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Day:", event.Day())

	fmt.Println("Title:", event.Title(), "Year:", event.Year(), "Month:", event.Month(), "Day:", event.Day())
}
