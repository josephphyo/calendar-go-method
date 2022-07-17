package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

// Get Method
func (e Event) Title() string {
	return e.title
}

// Set Method
func (e Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid Title")
	}
	e.title = title
	return nil
}
