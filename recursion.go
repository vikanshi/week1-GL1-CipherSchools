package main
import "fmt"

func main () {
    rec(5)
}

func rec(num int) {
    if num == 0 {
        return
    }
    if num%2 == 0 {
        rec(num - 1)
        fmt.Println(num-1)
    } else {
        rec(num - 1)
        fmt.Println(num-1)
    }
    fmt.Print(num - 1)
}