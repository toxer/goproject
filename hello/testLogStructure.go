package main 
import(
ls "github.com/toxer/goproject/logsutils"
"fmt"
)


func main() {
    l,_,_ := ls.CreateLogList("/tmp/log","server.log_","")
    fmt.Println(l)
}