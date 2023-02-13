package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	currentYear := time.Now().Year()
	var s string
	var age int
	fmt.Println("Enter your year of birth :")
	for {
		_, err := fmt.Scan(&s)
		age, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid year")
		} else {
			result := currentYear - age
			fmt.Println("You are", result, "years old")
			break
		}
	}
}
