//Pass by Value method

package main

import "fmt"

//function to swap 2 variables using pass by value method
func swap(x int, y int){
  x=x+y
  y=x-y
  x=x-y
  fmt.Printf("\nIn the swap function | X=%d,Y=%d.",x,y)
}

func main(){
  //declaring the variables
  x:=100
  y:=200
  fmt.Printf("Before calling the swap function | X=%d,Y=%d.",x,y)

  //calling function
  swap(x,y)

  //After the function call
  fmt.Printf("\nAfter calling the function | X=%d,Y=%d.",x,y)
}
