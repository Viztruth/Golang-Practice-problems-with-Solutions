//prime number or not
package main

import "fmt"

func check(a int)int{
  count:=0
  if a==1 || a==0{
    return 0
  }else{
  for i:=2;i<=a/2;i++{
    if a%i==0{
      count++
    }
  }
  if count>0{
    return 0
  }
  return 1;
}
  }

func main(){
  var a int
  fmt.Printf("Enter the number: ")
  fmt.Scanf("%v",&a)
  if check(a)==1{
    fmt.Printf("\n%v is a Prime number: ",a)
  }else{
    fmt.Printf("\n%v is not a Prime number: ",a)
  }
}
