/*
 * @file : 
 * Create: 2023.03.13
 * Author: ABdeeglr <abdeeglr@icloud.com>
 * Usage :
 */

package main

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 11
}

func main() {
	a, b, _ := vals()
	fmt.Println("a is", a)
	fmt.Println("b is", b)

	_, _, c := vals()
	fmt.Println("c is", c)
}