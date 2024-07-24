package time

import (
	"log"
	"time"
)

func WibToUTC(s string) (time.Time, error) {

	loc, err := time.LoadLocation(LOC_WIB)
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return time.Time{}, err
	}

	parsedTime, err := time.ParseInLocation(TIME_STAMP_WITH_SECOND, s, loc)
	if err != nil {
		log.Printf("Error parsing timestamp: %v", err)
		return time.Time{}, err
	}

	return parsedTime.UTC(), nil
}
