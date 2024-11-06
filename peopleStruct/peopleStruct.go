package peoplestruct

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	timeUti "github.com/vickz86/GoFunctions/timeUti"
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

// return a new People based on input
func CreatePeople(firstname, lastname string, birthDate time.Time) People {
	return People{Firstname: firstname, Lastname: lastname, Birth: birthDate}
}

// PrintAge calculates and prints the age of the person, returning it as an int
func (p People) PrintAge() int {
	// Get the current time
	currentTime := time.Now()

	// Calculate age based on the year difference
	age := currentTime.Year() - p.Birth.Year()

	// Adjust if the birth date has not yet occurred this year
	if currentTime.Month() < p.Birth.Month() || (currentTime.Month() == p.Birth.Month() && currentTime.Day() < p.Birth.Day()) {
		age--
	}

	// Print the age
	fmt.Printf("%s %s is %d years old\n", p.Firstname, p.Lastname, age)

	return age
}

// TODO
// func FormatBirthday () string*

// TimeToBirthday calculates the time remaining until the next birthday.
func (p People) TimeToBirthday() (time.Duration, int) {
	// Set the person's birthday to the current year
	nextBirthday := timeUti.ChangeYear(p.Birth, 0)

	// If birthday has already passed this year, set it to next year
	if !nextBirthday.After(time.Now()) {
		nextBirthday = timeUti.ChangeYear(nextBirthday, 1)
	}

	// also return the number of days until birthday
	timeUntil := time.Until(nextBirthday)

	// get the number of days until the next birthday
	nbDays := timeUti.DurationToDays(timeUntil)


	// Return the duration until the next birthday
	return timeUntil,nbDays
}
