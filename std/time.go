package std

import (
	"fmt"
	"time"
)

func TimeTest() {
	// time.Format
	now := time.Now()
	fmt.Println("Current time: ", now)
	fmt.Println("Format time to: ", now.Format("2006-01-02 15:04:05"))

	// time.Parse
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, err := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// time with location
	loc, err := time.LoadLocation("America/New_York")
	t2, err := time.ParseInLocation(longForm, "Feb 3, 2013 at 7:54pm (EST)", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parse with Location: ", t2)

	// convert to another time zone
	loc2, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	fmt.Println("Current time in LA: ", now.In(loc2))

	// time elspased
	duration := time.Since(t2)
	fmt.Printf("Elapsed %v since %v\n", duration, t2)

	// ticker
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := range 3 {
		<-ticker.C // waiting for the next ticker time
		fmt.Printf("ticker triggered %d times\n", i+1)
	}
}
