//greatest of 3 numbers
package main

import "fmt"

func main(){
  var a,b,c int
  fmt.Printf("\nEnter the 3 numbers a,b,c: ")
  fmt.Scanf("%v %v %v", &a,&b,&c);
  lar:=a
  if lar<b{
    lar=b
  }
  if lar<c{
    lar=c
  }
  fmt.Printf("\nThe largest no. is: %v", lar)
}
