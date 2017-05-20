package main

import "fmt"

func main() {
	// literal declaration of a map
	m := map[string]string{
		"keyOne":   "Keep Me",
		"keyTwo":   "Delete Me",
		"keyThree": "Not me",
	}

	fmt.Println("Map before deleting - ", m)
	fmt.Println("Deleting...")

	// To delete an entry use "delete(mapName , key-to-delete)"
	delete(m, "keyTwo")

	fmt.Println("Map after deleting - ", m)

}
