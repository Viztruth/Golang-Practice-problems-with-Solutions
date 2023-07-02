//check for anagram
package main

import(
  "fmt"
)

//function to perform bubblesort
func bubblesort(str []byte){
  n:=len(str)
  for i:=0;i<n-1;i++{
    for j:=0;j<n-1;j++{
      if str[j]>str[j+1]{
        temp:=str[j]
        str[j]=str[j+1]
        str[j+1]=temp
      }
    }
  }
}

//function to check for anagram
func checkanagram(str1, str2 string)int{
  //checking whether their lengths are equal
  if len(str1)==len(str2){
    //converting the string to slice of bytes as value of string cannot be changed once created
    s1 := []byte(str1)
    s2 := []byte(str2)
    bubblesort(s1)
    bubblesort(s2)
    for i:=0;i<len(s1);i++{
      if(s1[i]!=s2[i]){
        return 0
      }
    }
    return 1
    }
  return 0
}


func main(){
  str1:="anagaram"
  str2:="nagarama"
  fmt.Printf("The 2 strings are: %v, %v.",str1,str2)
  if checkanagram(str1,str2)==1{
    fmt.Println("\nThe two strings are anagram.")
  }else{
    fmt.Println("\nThe two strings are not anagram.")
  }
}
