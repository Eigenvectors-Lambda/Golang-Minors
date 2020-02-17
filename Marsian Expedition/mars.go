//Mars programme
package main

import (
	"fmt"
	"math/rand"
)


// main function :where it begins.
func main(){
	// personal Marsian data
	fmt.Printf("Marsian weight is:%v Newtons\n",63*9.81*0.38)
	fmt.Printf("Marsian Age is:%v years old.\n ",19*365/687)
	// How long to get to Mars?
	const lightSpeed=100800 //km/h
	var distance=96300000 //km
	fmt.Println("estimated time(SpaceX ITS):",distance/lightSpeed,"hours")
	// Random number generator
	var num= rand.Intn(12)+1
	fmt.Println(num)

	num=rand.Intn(12)+1
	fmt.Println(num)
	
}