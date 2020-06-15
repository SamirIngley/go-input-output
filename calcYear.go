package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

// Create a program that calculates the year, given the date of birth and age


// KEYBOARD IO
func keyboard() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter age: ")
    ioage, _ := reader.ReadString('\n')

    fmt.Print("Enter dob: ")
    iodob, _ := reader.ReadString('\n')

    // fmt.Print(ioage, iodb)
    fmt.Print(AgeYear(iodob, ioage))
    return AgeYear(iodob, ioage)
}


// Calculates year based on age of person
func AgeYear(dob time.Time, age time.Time) int{
    // Add the year of birth to current age to get the current year
    currentYear := dob.Year() + age.Year()
    return currentYear
}
