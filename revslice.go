//Reverse an slice

package main

import "fmt"

//function to reverse
func reverse(slc[]int, n int){
  for i:=0;i<n/2;i++{
    temp:=slc[i]
    slc[i]=slc[n-i-1]
    slc[n-i-1]=temp
  }
}


func main(){
  //declaring the slice
  slc:=[]int{1,2,3,4,5,6,7,8}
  n:=len(slc)
  
  fmt.Println("The slice: ",slc)
  reverse(slc,n)
  fmt.Println("\nThe Reversed slice: ",slc)
}
