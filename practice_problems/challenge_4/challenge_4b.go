package main

import (
	"fmt"
)

type Player struct {
    Name  string
    Score int
}

func main () {

	players := []Player{
		{Name: "Charlie", Score: 100},
		{Name: "Alice", Score: 100},
		{Name: "Bob", Score: 42},
	}


	winners, median, average := ComputeStats(players)
	
	fmt.Println("winners, median, average:", winners, median, average)

}

func ComputeStats(players []Player) (winners []string, median, average float64) {
    
	if len(players) == 0 {
		return nil, 0, 0
	}

	var totalScore int

	// Average
	for _, p := range players{
		totalScore += p.Score
	}

	average = float64(totalScore)/float64(len(players))



	var maxScore int = -1
	// Highest Score
	for _, p := range players {

		if p.Score > maxScore {
			maxScore = p.Score
			winners = []string{p.Name}
		} else if p.Score == maxScore {
			winners = append(winners, p.Name)
		}

	}

	scores := make([]int, len(players))

	for i, p := range players {
		scores[i] = p.Score
	}

	n := len(scores)

	for i := 0 ; i < n-1 ; i++ {
		for j:=0; j < n-i-1 ; j++ {
			if scores[j] > scores[j+1] {
				scores[j], scores[j+1] = scores[j+1], scores[j]
			}
		}
	}

	mid := n/2

	if n%2 == 0{
		median = float64(scores[mid])
	} else {
		median = float64(scores[mid-1]+scores[mid])/2.0
	}

	return winners, median, average

}