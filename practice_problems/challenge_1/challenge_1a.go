package main

import (
	"fmt"
	"strings"
)

type Passenger struct {
	BookingID string
	Name      string
	Seat      string
}

func main() {

	fmt.Println("Challenge #1a - ParsePassengerLine")

	line := "BK-123456|Ada Lovelace|12C"
	fmt.Println("result: ", ParsePassengerLine(line))
}

func ParsePassengerLine(line string) bool {
	
	fields := strings.Split(line, "|")
	fmt.Println(fields)

	if len(fields) != 3 {
		fmt.Println("length of fields more than three.")
		return false
	}

	bookingID := strings.TrimSpace(fields[0])
	fmt.Println("Booking ID:", bookingID)
	name := strings.TrimSpace(fields[1])
	fmt.Println("name:", name)
	seat := strings.TrimSpace(fields[2])
	fmt.Println("Seat:", seat)

	if bookingID == " " || name == " " || seat == " " {
		fmt.Println("fields are empty")
		return false
	}

	p := Passenger {
			BookingID: bookingID,
			Name:      name,
			Seat:      seat,
		}

	fmt.Println("Passenger:", p)

	return true
}
