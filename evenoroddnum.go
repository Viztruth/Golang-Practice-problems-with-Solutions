//Even or odd number
package main

import(
  "fmt"
)

func checkoff(a int)int{
  if(a%2==0){
    return 1;
  }
  return 0;
}

func main(){
  var n int
  fmt.Printf("Enter the number: ");
  fmt.Scanf("%v",&n);
  if checkoff(n)==1{
    fmt.Printf("\n%v is even!",n)
  }else{
    fmt.Printf("\n%v is odd!",n)
  }
}
