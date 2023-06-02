//Sum of all elements in slice

package main

import "fmt"

//function to find the sum
func sums(arr[]int,n int)int{
  sum:=0
  for i:=0;i<n;i++{
    sum=sum+arr[i]
  }
  return sum
}

func main(){
  //declaring the slice
  slc:=[]int{1,2,3,4,5,6,7,8,9}
  n:=len(slc)
  
  fmt.Println("The slice: ",slc)
  fmt.Printf("\nThe sum of all elements in slice: %d.",sums(slc,n))
}
