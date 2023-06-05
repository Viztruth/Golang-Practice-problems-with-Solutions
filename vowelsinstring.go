//counting number of vowels and consonants

package main

import(
  "fmt"
)


func main(){
  //declaring the string
  var name string="abdul narayan dsouza"
  fmt.Println("The string is: ",name)
  
  count:=0
  for i:=0;i<len(name);i++{
    if name[i]=='a' || name[i]=='e' || name[i]=='i' || name[i]=='o' || name[i]=='u'{
      count++
    }
  }
  fmt.Printf("\nThe no. of vowels and consonants is: %v, %v.",count,len(name)-count)
}
