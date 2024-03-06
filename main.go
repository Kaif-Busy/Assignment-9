package main

import (
	"fmt"
	"time"
)

func hasTimeComponent(date interface{}) bool {
	dateStr, ok := date.(string)
	if !ok {
		return false
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", dateStr)
	if err == nil {

		loc, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			return false
		}
		parsedTime = parsedTime.In(loc)

		return parsedTime.Hour() != 0 || parsedTime.Minute() != 0 || parsedTime.Second() != 0
	}

	_, err = time.Parse(time.RFC3339, dateStr)
	return err == nil

}



func main() {
	inputs := []interface{}{
		"2024-03-06 11:38:41.851892 +0000 UTC",
		"2024-03-06T21:30:00Z",
		"2024-03-06 18:30:00 +0000 UTC",
		"2024-03-06 00:00:00 +0000 UTC",
	}


	for _, input := range inputs {
		result := hasTimeComponent(input)
		fmt.Printf("Input: %v, Has time: %v\n", input, result)
	}
}
