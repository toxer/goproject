package main 
import(
ls "github.com/toxer/goproject/logstructure"

"fmt"
)


func main() {
	l,_ := ls.CreateLogList("/tmp/log","server.log_","")
	fmt.Println(l)
}