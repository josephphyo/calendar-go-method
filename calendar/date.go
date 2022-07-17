package calendar

import "errors"

// unexported fields
type Date struct {
	year  int
	month int
	day   int
}

// Get Methods
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

// Set Methods
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invaild Year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}
