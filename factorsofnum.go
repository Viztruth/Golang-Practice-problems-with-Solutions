// Program to generate factors of a number

package main
import(
    "fmt"
)

func genfactors(n int){
    for i:=1;i<=n/2;i++{
        if n%i==0{
            fmt.Printf("%v, ",i)
        }
    }
}

func main() {
    var a int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    fmt.Printf("\nThe factors are: ")
    genfactors(a)
}
