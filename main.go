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

	error = date.SetMonth(02)
	if error != nil {
		log.Fatal(error)
	}

	error = date.SetDay(14)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(date)
}
