//Checking Palindrome

package main

import(
  "fmt"
)


func main(){
  //declaring the string
  var name string="rotor"
  fmt.Println("The string is: ",name)

  //checking out whether the string's a palindrome
  leng:=len(name)
  count:=0
  for i:=0;i<leng/2;i++{
    if name[i]!=name[leng-i-1]{
      count++
    }
    }

  //Printing the result
  if count!=0{
    fmt.Println("\nThe string isn't a palindrome.")
  }else{
    fmt.Println("\nThe string is a palindrome.")
  }
}
