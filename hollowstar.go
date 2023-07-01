//hollow square pattern
package main

import(
  "fmt"
)


//function to generate hollow square pattern of side n
func pattern(n int) {
  //rows
  for i:=0;i<n;i++{
    //columns
    for j:=0;j<n;j++{

      //if first or last row, fill completely with *
      if i==0 || i==n-1{
        fmt.Printf("* ")
        }else{ 
        //if not, then fill only the edges with *
        if j==0 || j==n-1{
          fmt.Printf("* ")
        }else{
          fmt.Printf("  ")
        }
      }
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
