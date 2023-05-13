// Program to generate factorial

package main
import(
    "fmt"
)

func genfactorial(n int)int{
    facto:=1
    for i:=1;i<=n;i++{
        facto=facto*i
    }
    return facto
}

func main() {
    var a int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    fmt.Printf("\nThe factorial is: %v",genfactorial(a))
}
