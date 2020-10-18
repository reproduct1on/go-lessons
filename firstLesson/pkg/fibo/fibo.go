package fibo

import(

"fmt"

)



func calc(n int) int{
	if (n == 0){
		return 0
	} else if (n == 1)||(n == 2){
		return 1
	}else {
		return calc(n - 1) + calc(n - 2)
	}
}



func Solo(n int) {
	fmt.Println("fibonachi",calc(n))
}

func Multi(nums ...int){
	for _, n := range nums {
        fmt.Println("fibonachi myltiple",calc(n))
    }
}
