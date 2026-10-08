package main

import (
	"fmt"
)

type Passenger struct {
	BookingID string
	Name      string
	Seat      string
}

func main() {

	fmt.Println("Challenge #1a - ParsePassengerLine")

	line := "BK-123456|Ada Lovelace|12C"
	fmt.Println(ParsePassengerLine(line))
}

func ParsePassengerLine(line string) bool {
	return true
}
