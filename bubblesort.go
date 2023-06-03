//Sorting a slice using bubblesort

package main

import "fmt"

//bubblesort
func bubblesort(slc[]int, n int){
  for i:=0;i<n-1;i++{
    for j:=0;j<n-1;j++{
      if slc[j]>slc[j+1]{
        temp:=slc[j]
        slc[j]=slc[j+1]
        slc[j+1]=temp
      }
    }
  }
  }

func main(){
  //declaring the slice
  slc:=[]int{18,76,23,40,30,0,10,15}
  n:=len(slc)
  
  fmt.Println("The slice: ",slc)
  bubblesort(slc,n)
  fmt.Println("\nThe sorted slice: ",slc)
}
