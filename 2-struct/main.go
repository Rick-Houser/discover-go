package main

import "fmt"
import "time"

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func main() {
	form := "January 2, 2006"
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)
	t, err := time.Parse(form, u.DOB)
	result := time.Since(t).Seconds() / 31557600

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("%s was born in %s and would be %d years old today.\n", u.Name, u.City, int(result))
}
