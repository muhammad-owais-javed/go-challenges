package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Passenger struct {
	BookingID string
	Name      string
	Seat      string
}

func main() {

	fmt.Println("Challenge #1a - ParsePassengerLine")

	block := "BK-123456|Ada Lovelace|12C\n\nBK-000001|Ada Lovelace|12Z\nBK-654321|Grace Hopper|7A\nBK-123456|Ada L.|14F"

	passenger, invalidLines := ParsePassengerManifest(block)
	fmt.Println("result:", passenger, invalidLines)

}


func ParsePassengerManifest(block string) (map[string]Passenger, []int)  {

	manifest := make(map[string]Passenger)
	var invalidLines []int

	lines := strings.Split(block, "\n")
	fmt.Println("lines:", lines)


	for i, line := range lines {

		lineNum := i+1

		cleanedLine := strings.TrimSpace(line)
		
		if cleanedLine == ""{
			fmt.Println("EmptyLine at ",lineNum)
			continue
		}

		fields := strings.Split(line, "|")

		if len(fields) != 3 {
			fmt.Println("invalid line at ", lineNum)
			invalidLines = append(invalidLines, lineNum)
			continue
		}

		bookingID := strings.TrimSpace(fields[0])
		name := strings.TrimSpace(fields[1])
		seat := strings.TrimSpace(fields[2])

		// fmt.Println("Booking ID:", bookingID)
		// fmt.Println("name:", name)
		// fmt.Println("Seat:", seat)

		if bookingID == " " || name == " " || seat == " " {
			invalidLines = append(invalidLines, lineNum)
			fmt.Println("fields are empty at ", lineNum)
			continue
		}

		if seatValid(seat) != true {
			invalidLines = append(invalidLines, lineNum)
			fmt.Println("Seats are not valid at ", lineNum)
			continue
		}

		manifest[bookingID] = Passenger {
				BookingID : bookingID,
				Name    : name, 
				Seat    : seat,
			}
		

	}
	fmt.Println("manifest:", manifest)
	fmt.Println("invalidLines", invalidLines)
	return manifest, invalidLines

}


func seatValid(seat string) bool{

	if len(seat) < 2 {
		return false
	}

	lastChar := seat[len(seat)-1]
		if lastChar < 'A' || lastChar > 'F' {
		return false
	}

	rowStr := seat[:len(seat)-1]
	row, err := strconv.Atoi(rowStr)
		if err != nil {
		return false 
	}

	return row >= 1
}