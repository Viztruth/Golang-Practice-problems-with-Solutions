//2 ways of defining an array of size n
package main

import(
  "fmt"
)

func main(){
  //method 1 - hard-coding the elements
  arr1:=[...]int{10,9,8,7,6,5,4,3,2,1}
  n1:=len(arr1)
  fmt.Printf("The %d elements of array 1 are: ",n1);
  for i:=0;i<n1;i++{
    fmt.Printf("%d, ",arr1[i])
    }
  
  //method 2 - entering them manually
 var arr2 [6]int
  n2:=len(arr2)
  fmt.Printf("\nEnter the %d elements of array 2: ",n2)
  for i:=0;i<n2;i++{
    fmt.Scanf("%v",&arr2[i]);
  }
  fmt.Printf("The %d elements of array 2 are: ",n2)
  for i:=0;i<n2;i++{
    fmt.Printf("%d, ",arr2[i])
  }
}
