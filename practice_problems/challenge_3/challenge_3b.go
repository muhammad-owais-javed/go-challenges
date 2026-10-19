package main

import (
	"fmt"
	"time"
	"sort"
)

type ZonedSession struct {
    ID       string
    Start    string // local datetime
    End      string // local datetime
    Timezone string // "±HH:MM"
}

type ParsedSession struct {
	ID string
	Start time.Time
	End time.Time
}



func main(){

	sessions := []ZonedSession{
		{
			// Session A: 10:00 AM to 12:00 PM (UTC)
			ID:       "UTC-Morning",
			Start:    "2024-02-28 10:00",
			End:      "2024-02-28 12:00",
			Timezone: "+00:00",
		},
		{
			// Session B: 6:00 AM to 8:00 AM in New York (-05:00) -> 11:00 AM to 1:00 PM UTC (Conflicts with A)
			ID:       "NY-Early",
			Start:    "2024-02-28 06:00",
			End:      "2024-02-28 08:00",
			Timezone: "-05:00",
		},
		{
			// Session C: 12:00 PM to 2:00 PM (UTC) -> Starts when A ends (No conflict with A)
			ID:       "UTC-Midday",
			Start:    "2024-02-28 12:00",
			End:      "2024-02-28 14:00",
			Timezone: "+00:00",
		},
		{
			// Session D: 11:30 AM to 1:00 PM (UTC) -> Overlaps with others
			ID:       "UTC-Overlap",
			Start:    "2024-02-28 11:30",
			End:      "2024-02-28 13:00",
			Timezone: "+00:00",
		},
		{
			// Session E: INVALID (End time is before Start time) -> Should be ignored
			ID:       "Invalid-Reversed",
			Start:    "2024-02-28 15:00",
			End:      "2024-02-28 14:00",
			Timezone: "+00:00",
		},
	}

	
	fmt.Println("Output:", BuildConflictMap(sessions))
}

func BuildConflictMap(sessions []ZonedSession) map[string][]string {
    

	fmt.Println("Starting Processing Sessions:", sessions)
	conflicts := make(map[string][]string)
	var validSession []ParsedSession

	layout := "2006-01-02 15:04 -07:00"

	for _, s := range sessions {

		

		startStr := s.Start + " " + s.Timezone
		endStr := s.End + " " + s.Timezone

		startT , err := time.Parse(layout, startStr)
		if err != nil {
			fmt.Println("Invalid Session")
			continue
		}

		endT, err := time.Parse(layout, endStr)
		if err != nil {
			fmt.Println("Invalud Session")
			continue
		}

		if startT.After(endT) == true || endT.Before(startT) == true {
			continue
		}

		newSession := ParsedSession{
			ID: s.ID,
			Start: startT,
			End: endT,
		}

		validSession = append(validSession, newSession)

	}

	for i := 0 ; i < len(validSession) ; i++ {


		for j := 0; j < len(validSession); j++{

			sessionA := validSession[i]
			sessionB := validSession[j]

			if sessionA.Start.Before(sessionB.End) && sessionB.Start.Before(sessionA.End) {

				// Add B's ID to A's list
				conflicts[sessionA.ID] = append(conflicts[sessionA.ID], sessionB.ID)
				// Add A's ID to B's list
				conflicts[sessionB.ID] = append(conflicts[sessionB.ID], sessionA.ID)

			}


		}

	}

	for key := range conflicts {
		sort.Strings(conflicts[key])
	}

	return conflicts

}