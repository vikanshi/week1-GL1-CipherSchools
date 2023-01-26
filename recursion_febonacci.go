package main
import "fmt"

func main () {
fmt.Println(fib(3))
}

func fib (number int) int {
    //  base case
    if number == 0 || number == 1 {
        return number
    }
    
    
    result := fib(number-1) + fib(number-2)
    return result 
}