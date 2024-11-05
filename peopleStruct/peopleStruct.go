package peoplestruct

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Define the People struct
type People struct {
	Firstname string
	Lastname  string
	Birth     time.Time
}

// CreateBirthDate returns a time.Time from a string formatted as "DD-MM-YYYY"
func CreateBirthDate(birthString string) (time.Time, error) {

	// Split based on "-" separator
	elements := strings.Split(birthString, "-")
	if len(elements) != 3 {
		return time.Time{}, errors.New("error: incorrect birthString format, expected DD-MM-YYYY")
	}

	// Convert year
	yearInt, err := strconv.Atoi(elements[2])
	if err != nil {
		return time.Time{}, errors.New("error: invalid year supplied")
	}

	// Convert day
	dayInt, err := strconv.Atoi(elements[0])
	if err != nil {
		return time.Time{}, errors.New("error: invalid day supplied")
	}
	if dayInt < 1 || dayInt > 31 {
		return time.Time{}, errors.New("error: day must be between 1 and 31")
	}

	// Convert month
	monthInt, err := strconv.Atoi(elements[1])
	if err != nil {
		return time.Time{}, errors.New("error: invalid month supplied")
	}
	if monthInt < 1 || monthInt > 12 {
		return time.Time{}, errors.New("error: month must be between 1 and 12")
	}

	// Create the date, defaulting time to start of the day
	newDate := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)

	return newDate, nil
}
