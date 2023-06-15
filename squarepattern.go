//square pattern
package main

import(
  "fmt"
)


//function to generate square pattern of side n
func pattern(n int) {
  //rows
  for i:=0;i<n;i++{
    //columns
    for j:=0;j<n;j++{
      fmt.Printf("* ")
    }
    fmt.Printf("\n")
  }
}

func main(){
  //declaration
  var n int
  fmt.Print("Enter the side: ")
  fmt.Scanf("%v",&n)

  //output
  fmt.Printf("\nThe pattern - \n")
  pattern(n)
}
