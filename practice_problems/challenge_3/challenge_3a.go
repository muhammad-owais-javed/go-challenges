package main

import (
	"fmt"
	"time"
)

func main(){

	input := "2024-02-28 19:32"

	fmt.Println(FormatSessionTime(input))
}

func FormatSessionTime(input string) string {
    
	inputLayout := "2006-01-02 15:04"
	parsedTime, err := time.Parse(inputLayout, input)
	
	if err != nil {
		fmt.Println("Failed to Parse")
		return ""
	}
	
	outputLayout := "02 January 2006, 15:04PM"

	result := parsedTime.Format(outputLayout)

	return result
}