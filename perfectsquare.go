//perfect square
package main

import(
  "fmt"
  "math"
)

func main(){
  var n int
  fmt.Printf("Enter the number: ")
  fmt.Scanf("%v",&n)
  root:=math.Sqrt(float64(n))
  if root*root==float64(n){
    fmt.Printf("\n%v is a Perfect Square!",n)
  }else{
    fmt.Printf("\n%v is not a Perfect Square!",n)
  }
}
