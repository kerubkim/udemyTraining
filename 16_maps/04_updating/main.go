package main

import "fmt"

func main() {

	x := make(map[string]int)

	x["keyOne"] = 987
	x["keyTwo"] = 234
	x["keyThree"] = 909

	fmt.Println("Un updated - ", x)
	fmt.Println("Updating key two")

	x["keyTwo"] = 876132
	fmt.Println("Updated - ", x)

}
