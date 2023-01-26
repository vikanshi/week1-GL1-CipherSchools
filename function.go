package main
import "fmt"

func main() {
result, x := a()
fmt.Println(result)
fmt.Println(x)
result, x = c()
fmt.Println(result)
fmt.Println(x)
r, _ :=b(1,2,3,4,5,6,7,8,9)
fmt.Println(r)	
}
func a() (int, string) {
	return 122, "vikanshi"
}

func b(args ...int) (bool,int) {
    for _, v := range args {
        fmt.Print(v)
    }
    return true
}

func c() (i int, j string) {
	i =10
	j = "vikanshi"
	return
}