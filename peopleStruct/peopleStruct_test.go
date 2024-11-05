package peoplestruct

import (
	"testing"
	"time"
)

func TestCreateBirthDate(t *testing.T) {
	t.Parallel()

	// Define the expected date
	dateWant := time.Date(2019, time.September, 19, 0, 0, 0, 0, time.UTC)

	// Call CreateBirthDate with a valid date string
	dateGot, err := CreateBirthDate("19-09-2019")

	// Check if there was an unexpected error
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	// Compare the dates
	if !dateGot.Equal(dateWant) {
		t.Errorf("mismatched date: got %v, want %v", dateGot, dateWant)
	}
}