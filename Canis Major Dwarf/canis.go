package main

import (
	"fmt"
	"math/big"
)

func main(){
	lightSpeed:= big.NewInt(299792) // km s^-1
	secondsPerYear:= big.NewInt(3.154e7)

	canisDistance:= new(big.Int)
	canisDistance.SetString("236000000000000000",10)
	fmt.Println("Canis Major Dwarf is", canisDistance,"km away")

	seconds:=new(big.Int)
	seconds.Div(canisDistance,lightSpeed)

	years:= new(big.Int)
	years.Div(seconds,secondsPerYear)

	fmt.Println("That is", years, "light years away")
}

