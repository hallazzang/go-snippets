package main

import (
	"fmt"

	"github.com/hallazzang/read"
)

func CounterExample() {
	fmt.Print("input line> ")
	s, err := read.Line()
	if err != nil {
		panic(err)
	}

	// This is same as: counter := make(map[rune]int)
	counter := map[rune]int{}

	// And this will not work since the map is not initialized and
	// has a value of nil:
	// var counter map[rune]int

	for _, r := range s {
		// Note that the default value is 0 here, so there's no
		// need to check if there is such key in the map.
		counter[r]++
	}

	// The order of output is not guaranteed to be deterministic.
	for r, cnt := range counter {
		fmt.Printf("%c: %d time(s)\n", r, cnt)
	}
}

func SearchExample() {
	// users is a map from username to actual user information.
	users := map[string]struct {
		Name string
		Age  int
	}{
		"john": {"John Doe", 27},
		"jane": {"Jane Doe", 25},
		"joe":  {"Joe Schmoe", 30},
	}

	fmt.Print("search user by username> ")
	username, err := read.String()
	if err != nil {
		panic(err)
	}

	user, ok := users[username]
	if !ok {
		fmt.Println("there is no such user")
	} else {
		fmt.Printf("Name: %s, age: %d\n", user.Name, user.Age)
	}
}

func FindDuplicationExample() {
	// For simple duplication check, we can use struct{} for value type.
	// In these cases, a map acts like a set.
	set := map[int]struct{}{}

	numbers := []int{10, 2, 50, 7, 8, 40, 20, 8, 4}

	found := false
	for i, n := range numbers {
		if _, ok := set[n]; ok {
			fmt.Printf("found duplication at index %d!\n", i)
			found = true
			break
		}
		set[n] = struct{}{}
	}
	if !found {
		fmt.Println("duplication not found")
	}
}

func main() {
	fmt.Println("<Counter example>")
	CounterExample()

	fmt.Println("<Search example>")
	SearchExample()

	fmt.Println("<Find duplication example>")
	FindDuplicationExample()
}
