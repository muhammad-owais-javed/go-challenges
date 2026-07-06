package main

import (
	"fmt"
)

type User struct{
	Name string
}

type Greeter interface{
	Greet() string
}

func (u User) Greet() string{
	return "Hello from" + u.Name
}

func main() {
	
	fmt.Println("Hello World!")

	// Primitive Data Types
	
	var isCloudNative bool = true	// True or False
	var developerName string = "Owais" // Text

	var userCount int = 1000
	var loopCounter uint = 999 // Must be positive

	var audioVolumeDb int8 = -12 // Small range (-128 to 127).
	var cityPopulation int64 = 1000000000 // Large range (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
	var starsInTheSky uint64 = 10000000000 // Large range (0 to 18,446,744,073,709,551,615)

	var rgbColor uint32 = 0xFF0000 // Hexadecimal representation of a color (Red)
	var gameTextureResolution uint16 = 300 // Small range (0 to 65,535)
	var nationalDebt uint32 = 1000000000 // Large range (0 to 4,294,967,295)
	var satelliteDistance uint64 = 30000000000000 // Large range (0 to 18,446,744,073,709,551,615)
	
	var productRating float32 = 4.5 // Decimal number with single precision
	var piValue float64 = 3.141592653589793 // Decimal number with double precision

	var asciiChar byte = 'A' // Alias for uint8, Represents a single character (ASCII)
	var unicodeChar rune = '世' // Alias for int32,Represents a single character (Unicode)

	fmt.Printf("%-25s | %-15s | %s\n", "Variable Name", "Go Type", "Value")
	fmt.Println("---------------------------------------------------------------------")

	fmt.Printf(
		"%-25s | %-15s | %t\n"+
			"%-25s | %-15s | %s\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | 0x%X\n"+ // Printed as Hexadecimal
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %d\n"+
			"%-25s | %-15s | %.1f\n"+ // Fixed to 1 decimal place
			"%-25s | %-15s | %.15f\n"+ // Fixed to 15 decimal places
			"%-25s | %-15s | %c (numeric: %d)\n"+ // Printed as char AND raw byte
			"%-25s | %-15s | %c (numeric: %d)\n", // Printed as char AND raw rune

		"isCloudNative", "bool", isCloudNative,
		"developerName", "string", developerName,
		"userCount", "int", userCount,
		"loopCounter", "uint", loopCounter,
		"audioVolumeDb", "int8", audioVolumeDb,
		"cityPopulation", "int64", cityPopulation,
		"starsInTheSky", "uint64", starsInTheSky,
		"rgbColor", "uint32", rgbColor,
		"gameTextureResolution", "uint16", gameTextureResolution,
		"nationalDebt", "uint32", nationalDebt,
		"satelliteDistance", "uint64", satelliteDistance,
		"productRating", "float32", productRating,
		"piValue", "float64", piValue,
		"asciiChar", "byte (uint8)", asciiChar, asciiChar,
		"unicodeChar", "rune (int32)", unicodeChar, unicodeChar,
	)


	// Composite/Reference Data Types

	var num int = 65
	var ptrAddr *int = &num

	var fruits [3]string
	fruits = [3]string{"Apple", "Oranges", "Banana"}
	var ports [3]int = [3]int{80, 8080, 443}


	fmt.Printf("\n%d, %s, %d\n", ptrAddr, fruits, ports)

}

