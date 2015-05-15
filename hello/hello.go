package main
import ("fmt"
//"encoding/json"
"os"


)
func main() {
	path := "/tmp/data.txt"
	fi, err := os.Stat(path)
	if err != nil {
		return
	}
    //converto la stat del file in stringa
    str:=fmt.Sprintln(fi.Sys())
	fmt.Println(str)


} 
