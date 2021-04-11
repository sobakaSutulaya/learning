package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	slice := createSlice()

	for _, value := range slice {
		fmt.Printf("%v is : ", value)
		if (value % 2 == 0) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}

func createSlice() []int {
	slice := []int{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < 10; i++ {
		slice = append(slice, r.Int())
	}
	return slice
}
