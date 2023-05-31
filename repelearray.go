//replacing an element in slice
package main

import(
  "fmt"
)

func main(){
  //declaring variables
  var len,ele,loc int
  
  //defining length of slice
  fmt.Printf("Enter the length: ");
  fmt.Scanf("%d %d",&len);

  //declaring a slice
  slc:=make([]int, len); 

  //entering the elements
  fmt.Printf("\nEnter the elements: ")
  for i:=0;i<len;i++{
    fmt.Scanf("%d",&slc[i])
  }

  //Input the element and loc
  fmt.Printf("\nEnter the location and element(0-%d): ",len-1)
  fmt.Scanf("%d %d",&loc,&ele)

  //replacing the element at loc
  if loc>len || loc<0 {
    fmt.Println("\nInvalid Input!")
  }else{
    fmt.Println(slc[loc]," has been successfully replaced by ",ele,"!")
    slc[loc]=ele;
      }

  //displaying the modified slice
  fmt.Printf("\nThe slice after replacement are: ")
  for i:=0;i<len;i++{
    fmt.Printf("%d, ",slc[i])
  }
  }
