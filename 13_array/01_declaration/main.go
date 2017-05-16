package main

import "fmt"

func main() {
	// declaring an array with " [ number ]" number is how many items will be in the array
	// if it just " [] " without the number inside, it is a SLICE
	// array is not dynamic (change in length)
	var arr [58]string // declaring
	fmt.Println(arr)
	fmt.Println(len(arr))        // checking the length of an array with len()
	fmt.Println(arr[6])          // accessing array of index 6 with is the 7th item
	for i := 65; i <= 122; i++ { // starting at 65 because in ASCII 65 is letter A
		arr[i-65] = string(i) // converts the int to an ASCII letter
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[6])

	fmt.Println("-------INT--------")
	var intArr [12]int
	for i := 0; i < len(intArr); i++ {
		intArr[i+0] = i
	}
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(intArr[6])
}
