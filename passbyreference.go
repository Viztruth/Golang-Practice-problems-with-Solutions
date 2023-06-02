//Pass by reference method

package main

import "fmt"

//function to swap 2 variables using pass by reference method
func swap(x *int, y *int){
  *x=*x+*y;
  *y=*x-*y;
  *x=*x-*y
}

func main(){
  //declaring the varicbles
  x:=100
  y:=200
  fmt.Printf("Before swapping | X=%d,Y=%d.",x,y)

  //calling function
  swap(&x,&y)
  fmt.Printf("\nAfter swapping | X=%d,Y=%d.",x,y)
}
