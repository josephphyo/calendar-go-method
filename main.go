package main

import (
	"fmt"
	"log"

	"github.com/josephphyo/export-packages-test/setter-methods/calendar"
)

func main() {
	date := calendar.Date{}

	error := date.SetYear(2022)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Year:", date.Year())

	error = date.SetMonth(02)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Month:", date.Month())

	error = date.SetDay(14)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Day:", date.Day())

	fmt.Println("Year:", date.Year(), "Month:", date.Month(), "Day:", date.Day())
}
