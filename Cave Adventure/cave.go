package main
import (
	"fmt"
	"strings"
)
func main(){
	// exit outputs Boolean when checking the command for the term "outside" 
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command="walk outside"
	var exit=strings.Contains(command,"outside")

	fmt.Println("You leave the cave:",exit)

	// Comparing to output Boolean
	fmt.Println("There is a sign near the entrance that reads'No Minors'.")
	var age=19
	var minor= age<18 // The Boolean for minors
	fmt.Printf("At age %v, am I a minor? %v\n",age,minor)

	// Branching Paths with if-Conditionals
	command="go east"
	if command== "go east"{
		fmt.Println("You head further up the mountain.")
	}else if command=="go inside"{
		fmt.Println("You enter the cave where you live out the rest of your life.")
	}else{
		fmt.Println("Didn't quite get that.")
	}
	// Checks future years if they are leap years with logical operators
	fmt.Println("The year is 2100, should you leap?")
	var year =2100
	var leap=year%400==0||(year%4==0 && year%100 !=0)
	if leap{
		fmt.Println("Look before you leap!")
	}else{
		fmt.Println("Keep your feet on the ground.")
	}
	var haveTorch= true
	var litTorch= false

	if !haveTorch||!litTorch{
		fmt.Println("Too dark! Nothing to see here.")
	}

	// Branching paths with Switch
	fmt.Println("There is a cavern entrance here and a path to the east.")
	command="go inside"

	switch command{
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave","go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors!'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
	// Entering rooms with Switch and Conditionals
	var room= "lake"

	switch{
	case room=="cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room=="lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room=="underwater":
		fmt.Println("The water is freezing cold.")
	}
}