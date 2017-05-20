package main

import "fmt"

func main() {

	x := make(map[string]string)
	x["keyOne"] = "Yooo"
	x["keyTwo"] = "Supp"

	fmt.Println(x)
	fmt.Println("The length of map x is : ", len(x))
}
