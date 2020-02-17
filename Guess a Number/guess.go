package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var (
		userNumber=75
		guessNumber=rand.Intn(100)+1 // !! REMEMBER THE "rand.Intn(x)" only counts from 0 to (x-1) but with a +1 it counts from 1 to x !!
	)
	for{
		if guessNumber<userNumber{
			fmt.Println(guessNumber,", Too Small")
		}else if guessNumber>userNumber{
			fmt.Println(guessNumber,", Too Big")
		}else{
			fmt.Println(guessNumber,", it is the same")
			break
		}
		guessNumber=rand.Intn(100)+1
	}

}
