package fibo
import(

	"testing"

)

func TestFibo(t  *testing.T){
	testArray:=[]int{0,1,1,2,3,5,8,13,21,34,55,}
	for i:= 0 ; i < 11; i++{
		value := calc(i)
		if (value != testArray[i]){
			t.Error("Expected ",testArray[i],", got ", value, "at ",i)
		}
	}
	
}