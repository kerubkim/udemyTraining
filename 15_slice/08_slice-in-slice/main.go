package main

import "fmt"

func main() {

	// the depth of multi-dimensional slice is how many " [] " you add ex. " [][][][][]string "
	allStudents := make([][]string, 0) // declaring a two dimensional slice that holds - a slice of strings

	studentOne := make([]string, 3)
	studentOne[0] = "Kerub"
	studentOne[1] = "Querubin"
	studentOne[2] = "05-04-1995"

	studentTwo := make([]string, 3)
	studentTwo[0] = "k"
	studentTwo[1] = "kim"
	studentTwo[2] = "Student"

	allStudents = append(allStudents, studentOne)

	allStudents = append(allStudents, studentTwo)

	fmt.Println(allStudents)
}
