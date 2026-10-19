package main

import (
	"fmt"
)

func main(){



}


func IsValidParkingLot(parking string) bool {

	lines := strings.Split(parking, "\n")
	
	hasGate := false
	hasSpace := false

	for r := 0; r < len(lines); r++ {

		for c := 0; c < len(lines[r]); c++ {
			char := lines[r][c]

			if char != 'E' && char != 'O' && char != 'W' && char != 'G' {
				return false
			}

			if char == 'E' || char == 'O' {
				hasSpace = true
			}
			if char == 'G' {
				hasGate = true
			}

			isEdge := false
			
			if r == 0 || r == len(lines)-1 || c == 0 || c == len(lines[r])-1 {
				isEdge = true
			} else if c >= len(lines[r-1]) || c >= len(lines[r+1]) {
				isEdge = true
			}

			if isEdge {
				if char != 'W' && char != 'G' {
					return false
				}
			} else {
				if char == 'G' {
					return false
				}
			}
		}
	}

	return hasGate && hasSpace
}