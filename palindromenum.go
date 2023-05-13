// Program to check for palindrome

package main
import "fmt"

func main() {
    var a,rev int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    rev=0;
    cpy:=a
    for a>0{
        rev=rev*10+a%10
        a=a/10
    }
    fmt.Printf("The number reversed: %v",rev)
    if rev==cpy{
        fmt.Printf("\nThe number is a palindrome.")
    }else{
        fmt.Printf("\nThe number is not a palindrome.")
    }
}
