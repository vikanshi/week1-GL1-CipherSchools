package main
import "fmt"

func main() {
    
    number := 10
    fmt.Println(number)
    
    getInt := func(x int) int {
        fmt.Println("In a function")
        return 10 + x 
    } 
    getInt(19)
    
     j := func(x int) int {
        fmt.Println("In a 1st function")
        return 20 + x 
    } (19)
    
    fmt.Println(j)
    
    g(getInt)
}

func g (getInt func(int) int) {
    getInt(78)
}
