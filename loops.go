package main
import "fmt"

func main() {

	for i := 0 ; i<5 ; i++ {
		for j := 0 ; j<5 ; j++ {
			fmt.Println(i,j)
	} 
}

for i := 0; i <3 ; i++ {
	if i==1 {
		 continue
	}
	fmt.Println(i)
}

nums := []int{1,2,3,4,5}
for k, v:=range nums {
	fmt.Println(k)
	fmt.Println(v)
}
}