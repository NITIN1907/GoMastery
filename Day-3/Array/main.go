package main

import "fmt"

func main() {
	fmt.Println("Arrays....")
	//array creation...
	var num [5]int
	num[1] = 1
	num[2] = 2
	num[3] = 3
	fmt.Println(num)

	//array init
	var data = [4]int{1, 2, 3, 4}
	fmt.Println(data)
	fmt.Printf("data is: %d\n", data)

	//length of array
	fmt.Println(len(data))

	//getting data from index
	fmt.Println(data[1])

	// Without specifying size (compiler counts it)
	var count = [...]int{5, 3, 2, 1, 3}
	fmt.Println(count)
	fmt.Println(len(count))

	//using loops
	for i := 0; i < len(count); i++ {
		fmt.Println(count[i])
	}

	//range loop
	for index, value := range data {
		fmt.Println("Index: ", index, " value: ", value)
	}
	//make() is a built-in function used to create slices, maps, and channels with allocated memory — especially when you want to define the size or capacity ahead of time
	//? make(type, length, capacity).
	//type: the type of slice, map, or channel
	//length: how many elements it should initially have
	//capacity: (optional, only for slices) total space allocated

	s := make([]int, 3, 5)
	fmt.Println(s)      // Output: [0 0 0]
	fmt.Println(len(s)) // Output: 3
	fmt.Println(cap(s)) // Output: 5
}
