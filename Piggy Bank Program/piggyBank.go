package main

import (
	"fmt"
	"math/rand"
)

func main(){
	nickel:=0.05
	dimes:=0.10
	quarters:=0.25
	piggyBank:=0.0
	for piggyBank<20.00{
		randomChoice:=rand.Intn(3)
		switch randomChoice{
		case 0:
			piggyBank+=nickel
		case 1:
			piggyBank+=dimes
		case 2:
			piggyBank+=quarters 
		}
		fmt.Printf("Money in the piggy bank is $ %4.2f \n",piggyBank)
	}
	
}





