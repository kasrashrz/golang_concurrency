package main 

import(
	"time"
	"fmt"
)

func main(){
	go count("kasra")
	count("asghar")
}

func count(str string){
	for i := 0; i <= 10; i++{
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}