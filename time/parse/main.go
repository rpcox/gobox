package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	mytime := "2025-04-03 19:15:00"
	// Parsing a time string without a time zone
	timeWithoutZone, err := time.Parse(layout, mytime)
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}
	fmt.Println("Parsing time without time zone:", mytime, ` => `, timeWithoutZone)
	fmt.Println("Location of parsed time defaults to:", timeWithoutZone.Location())

	// Parsing a time string with a UTC time zone
	timeWithUTC, err := time.Parse(layout+` MST`, mytime+` UTC`)
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}
	fmt.Println("Parsing time using time zone:", layout+` MST`, ` => `, timeWithUTC)
	fmt.Println("Location converts to:", timeWithUTC.Location())

	// Parsing a time string with a specific time zone offset
	timeWithOffset, err := time.Parse(layout+` -0700 MST`, mytime+` -0500 EST`)
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}
	fmt.Println("Parsing time using time zone offset:", layout+` -0700 MST`, ` => `, timeWithOffset)
	fmt.Println("Location is now:", timeWithOffset.Location())

	// Parsing a time string with a specific time zone offset
	timeWithJustOffset, err := time.Parse(layout+` -0700`, mytime+` -0500`)
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}
	fmt.Println("Parsing time using just the zone offset:", layout+` -0700`, ` => `, timeWithJustOffset)
	fmt.Println("Location is now:", timeWithOffset.Location())
}
