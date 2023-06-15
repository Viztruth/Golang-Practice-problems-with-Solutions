//sum of digits in a string
package main

import(
  "fmt"
)


//Function sum of digits in a string
func sumofnums(name string)int {
  sum:=0
  leng:=len(name)
  for i:=leng-1;i>=0;i--{
    if name[i]=='1'{
      sum=sum+1
    }
    if name[i]=='2'{
      sum=sum+2
    }
    if name[i]=='3'{
      sum=sum+3
    }
    if name[i]=='4'{
      sum=sum+4
    }
    if name[i]=='5'{
      sum=sum+5
    }
    if name[i]=='6'{
      sum=sum+6
    }
    if name[i]=='7'{
      sum=sum+7
    }
    if name[i]=='8'{
      sum=sum+8
    }
    if name[i]=='9'{
      sum=sum+9
    }
    }
  return sum
}

func main(){
  //declaring the string
  var name string="#1Classics of 1900s"
  fmt.Println("The string is: ",name)

  //passing a string to function and display output
  fmt.Printf("\nThe sum is: %v",sumofnums(name))
}
