package main

import (
	"fmt"
	"log"

	"github.com/josephphyo/export-packages-test/setter-methods/calendar"
)

func main() {
	date := calendar.Date{}

	error := date.SetYear(2019)
	if error != nil {
		log.Fatal(error)
	}

	error = date.SetMonth(10)
	if error != nil {
		log.Fatal(error)
	}

	error = date.SetDay(12)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Year:", date.year, "Month:", date.month, "Day:", date.day)
}
