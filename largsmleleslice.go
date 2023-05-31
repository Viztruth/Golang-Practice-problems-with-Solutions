//Largest, smallest element in slice

package main

import "fmt"

func main(){

  //defining a slice
  slc:=[]int{5,3,8,3,0,8,39,56,31}
  n:=len(slc);

  //assuming the largest and the smallest are the first elements of the slice
  l:=slc[0]
  s:=slc[0]

  //finding the largest and the smallest
  for i:=0;i<n;i++{
    if l<slc[i]{
      l=slc[i]
    }

    if s>slc[i]{
      s=slc[i]
    }
  }

  //displying the output
  fmt.Printf("\nThe largest and the smallest are: %d, %d",l,s)
}
