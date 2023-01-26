package main
import "fmt"

func main () {
    
fmt.Println(fact(3))
}
func fact (number int ) int {
    
    // base case
    if number == 1 || number == 0 {
        return 1
    }
    // corner case
    if number < 0 {
        return fmt.Println("Invalid Number")
    }
     result := number * fact(number-1)
     return result
    
}