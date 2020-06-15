import "time"

// Create a program that calculates the year, given the date of birth and age

// Calculates year based on age of person
func AgeYear(dob time.Time, age time.Time) int{
    // Add the year of birth to current age to get the current year
    currentYear := dob.Year() + age.Year()
    return currentYear
}
