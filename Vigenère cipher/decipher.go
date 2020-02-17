package main

import(
	"fmt"
)

func main(){
	cipherText:="CSOITEUIWUIZNSROCNKFD"
	keyword:="GOLANG"
	keyIndex:=0
	decode:=""
	
	for _,c:=range cipherText{
		cString:=string(c)
		cByte:=cString[0]
		decode+=string((cByte-keyword[keyIndex]+26)%26+'A')
		keyIndex++
		keyIndex%=len(keyword)
		

		
	}
	fmt.Printf("%v",decode)
}
