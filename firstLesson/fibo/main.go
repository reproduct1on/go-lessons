package main
import(

	"lesson1/pkg/fibo"
	"flag"
	"fmt"
)

func main(){
	
	var nFlag = flag.Int("n", 10, "help message for flag n")
	flag.Parse()
	if *nFlag > 19 {
		fmt.Println("to big value")
		return
	}
	fibo.Solo(*nFlag)
	fibo.Multi(4,3,4,2,10)


}


