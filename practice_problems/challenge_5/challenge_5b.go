package main

import (
	"fmt"
	"strings"
)

type Boarding struct {
    Name string
    Seat string
    Gate string
}



func main(){

	items := []Boarding{
		{Name: " Ada ", Seat: " 12C ", Gate: " A1 "},
		{Name: "Bob", Seat: "", Gate: "A2"},
		{Name: "Cara", Seat: "7A", Gate: "B3"},
		}

	fmt.Printf("\n%s\n",RenderBoardingLines(items))
	
}


func RenderBoardingLines(items []Boarding) []string {
    
	var validitems []Boarding


	maxNameLen := 0
	maxSeatLen := 0
	maxGateLen := 0

	for _, item := range items {

		name := strings.TrimSpace(item.Name)
		seat := strings.TrimSpace(item.Seat)
		gate := strings.TrimSpace(item.Gate)

		if name == "" || seat == "" || gate == "" {
			fmt.Println("name or seat or gate is empty")
			continue
		}

		cleanBoarding := Boarding{
			Name: name,
			Gate: gate,
			Seat: seat,
		}

		validitems = append(validitems, cleanBoarding)

		if len(name) > maxNameLen {
			maxNameLen = len(name)
		}
		
		if len(seat) > maxSeatLen {
			maxSeatLen = len(seat)
		}
		
		if len(gate) > maxGateLen {
			maxGateLen = len(gate)
		}
	}

	fmt.Println("validitems:", validitems)

	var result []string


	for _, item := range validitems {

		formattedLine := fmt.Sprintf("%-*s | Seat %-*s | Gate %-*s",
			maxNameLen, item.Name,
			maxSeatLen, item.Seat,
			maxGateLen, item.Gate,
		)

		result = append(result, formattedLine)
	}

	return result
}