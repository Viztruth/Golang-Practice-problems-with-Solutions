//Declaring a string and converting to uppercase & lowercase

package main

import(
  "fmt"
  "strings"
)


func main(){
  //declaring the string
  var name string="Abdul narayan dsouza"
  fmt.Println("The string is: ",name)
  
  // to uppercase
  fmt.Println("The string in uppercase: ",strings.ToUpper(name))
  
  //to lowercase
  fmt.Println("The string in lowercase: ",strings.ToLower(name))

  //to titlecase
  fmt.Println("The string in titlecase: ",strings.Title(name))
}
