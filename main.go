package main

import (
	"fmt"
)

func main() {
	friends := []string{
		"Budi",
		"Rudi",
		"Fajar",
		"Indah",
		"Ridwan",
		"Risman",
		"Rian",
		"Kosasih",
		"Kurnia",
		"Widya",
	}

	for k, friend := range friends {
		fmt.Printf("%d. %s", k+1, friend+" \r\n")
	}
}
