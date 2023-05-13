// Program to check for armstrong number

package main
import(
    "fmt"
    "math"
)

func order(n int)int{
    count:=0
    for n>0{
        count++
        n=n/10
    }
    return count
}

func checkarmsrtong(n int) bool{
    sums:=0
    cpy:=n
    ord:=order(n)
    for n>0{
        sums=sums+int(math.Pow(float64(n%10),float64(ord)))
        n=n/10
    }
    if sums==cpy{
        return true
    }else{
        return false
    }
}

func main() {
    var a int
    fmt.Println("Enter the number: ")
    fmt.Scanf("%v",&a);
    if checkarmsrtong(a){
        fmt.Printf("\nIt's an armstrong number.")
    }else{
        fmt.Printf("\nIt's not an armstrong number.")
    }
}
