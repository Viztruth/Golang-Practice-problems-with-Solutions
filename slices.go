//2 ways of defining slices of size n
package main

import(
  "fmt"
)

func main(){
  //method 1 - hard-coding the elements
  slc1:=[]int{10,9,8,7,6,5,4,3,2,1}
  n1:=len(slc1)
  fmt.Printf("The %d elements of slice 1 are: ",n1);
  for i:=0;i<n1;i++{
    fmt.Printf("%d, ",slc1[i])
    }
  
  //method 2 - entering them manually
  var len, cap int;
  fmt.Printf("\n\nEnter the length and capacity: "); //remember len<=cap
  fmt.Scanf("%d %d",&len,&cap);
  slc2:=make([]int,len,cap)
  fmt.Print("Enter the elements: ")
  for i:=0;i<len;i++{
    fmt.Scanf("%d",&slc2[i])
  }
  fmt.Printf("The elements of slice 2 are: ")
  for i:=0;i<len;i++{
    fmt.Printf("%d, ",slc2[i])
  }
}
