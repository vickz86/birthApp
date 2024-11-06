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


func TestPrintAge(t *testing.T) {
	Victor := People{
		Firstname: "victor",
		Lastname: "Chassaigne",
		Birth: time.Date(1986,11,9,0,0,0,0,time.UTC),
	}
	
	yearWant := 37
	
	yearGot := Victor.PrintAge()
	
	if yearWant!=yearGot{
		t.Errorf("wanted %v amd got %v",yearWant,yearGot)
	}
	
}

func TestTimeToBirthday(t *testing.T) {
	Victor := People{
		Firstname: "victor",
		Lastname: "Chassaigne",
		Birth: time.Date(1986,11,9,0,0,0,0,time.UTC),
	}
	
	_, gotDays := Victor.TimeToBirthday()
	wantDays := 2
	
	if wantDays != gotDays{
		t.Errorf("wanted %v amd got %v",wantDays,gotDays)
	}
	
	
}
