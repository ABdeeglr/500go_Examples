/*
 * @file : 
 * Create: 2023.03.13
 * Author: ABdeeglr <abdeeglr@icloud.com>
 * Usage :
 */

package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}