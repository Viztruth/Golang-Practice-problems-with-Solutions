package main

import(
  "fmt"
)


func main(){
  //declaring the string
  var name string="Abdul narayan dsouza"
  fmt.Println("The string is: ",name)

  //printing string in reverse order
  fmt.Printf("\nThe string in reverse order: ")
  leng:=len(name)
  for i:=leng-1;i>=0;i--{
    fmt.Print(string(name[i]))
    }
}
