// Program to generate fibonacci series

package main
import(
    "fmt"
)

func genfibbonacci(n int){
    a:=0
    b:=1
    fmt.Printf("\nThe fibonacci series are: %v, %v, ",a,b)
    for i:=0;i<n-2;i++{
        sum:=a+b
        a=b
        b=sum
        fmt.Printf("%v, ",b)
    }
}

func main() {
    var a int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    genfibbonacci(a)
}
