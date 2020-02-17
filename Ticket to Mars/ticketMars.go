package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Printf("%-17v %-5v %-11v %-5v \n","Spaceline","Days","Trip type","Price")
	fmt.Println("=========================================")
	for count:=0;count<=19;count++{
			//this produces a random number from 1 to 4 in order to pick the Spaceline
	    randomNumber:=rand.Intn(4)+1
	    spaceline:=""
	    switch randomNumber {
	    case 1:
		    spaceline="Virgin Galactic"
	    case 2:
		    spaceline="SpaceX"
	    case 3:
		    spaceline="NASA"
	    case 4:
		    spaceline="JAXA"
	    }
	    speed:= rand.Intn(14)+16 //km/s
	    distance2Mars:= 62100000 //km
	    days:=(distance2Mars/speed)/(3600*24)
	    price:= (speed+20)*2 //millions USD
	    randomTrip:=rand.Intn(2)
	    tripType:=""
	    switch randomTrip{
	    case 0:
		    tripType="One-way"
	    case 1:
		    tripType="Round-trip"
		    price*=2
	    }
	    fmt.Printf("%-17v %5v %-11v %v %4v \n",spaceline,days,tripType,"$",price)
	}
	


}
