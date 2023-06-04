//frequency of elements

package main

import "fmt"

//function to print out frequency
func checkfreq(slc[]int, n int){
  //slice to check for repetition
  maps := []int{}
  siz:=0
  
  for i:=0;i<n;i++{
    freq:=1
    count:=0
    //check whether the element is in maps
    for j:=0;j<siz;j++{
      if maps[j]==slc[i]{
        count++
      }
      }

      //if it's not already present in map, add it to maps and find its freq
      if count==0{
        maps=append(maps,slc[i])
        siz++
        for j:=i+1;j<n;j++{
          if slc[i]==slc[j]{
            freq++
          }
        }
        fmt.Printf("\nThe frequency of %d is %d.",slc[i],freq)
      }
  }
}

func main(){
  //declaring the slices
  slc:=[]int{18,76,23,40,52,30,0,10,80,52,40,52,58}
  n:=len(slc)

  fmt.Println("The slice: ",slc)
  checkfreq(slc,n)
}
