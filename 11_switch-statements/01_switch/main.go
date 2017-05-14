package main

import "fmt"

func main() {

	switch "Kerub" { // checks if "Kerub" is one of the case
	case "Kim":
		fmt.Println("Hello Kim")
		break // this is not needed on GO
	case "Kerub":
		fmt.Println("Whats good kerub")
	case "Querubin":
		fmt.Println("okok querubin")
	default:
		fmt.Println("Who are you?")
	}
	// no default fallthrough
	// fallthrough is optional
	//	-- you can specify fallthrough by explicitly stating it
	//	-- break isn't needed like in other languages
}
