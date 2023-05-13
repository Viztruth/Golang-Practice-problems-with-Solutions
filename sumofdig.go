// Program to print sum of digits of number

package main
import "fmt"

func main() {
    var a,sum int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    sum=0
    for a>0{
        sum=sum+a%10
        a=a/10
    }
    fmt.Printf("\nThe sum is %v",sum)
}
