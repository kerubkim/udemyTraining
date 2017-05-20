package main

import "fmt"

func main() {

	// Declaring a map
	map1 := make(map[string]int) // declaring a map with a key of strings and values of ints
	map1["keyone"] = 123
	map1["keytwo"] = 456

	fmt.Println(map1)
	fmt.Println("The value with the key 'keyone' is : ", map1["keyone"]) // printing one key

	// A Composite Literal way of declaring a map
	fmt.Println("============================")
	map2 := map[string]int{"foo": 987, "bar": 5432}
	fmt.Println(map2)
	fmt.Println("The value with the key of 'bar' is : ", map2["bar"])

}
