package main
import ("fmt"
"os"


)
func main() {
    path := "/tmp/data.txt"
    fi, err := os.Stat(path)
    if err != nil {
        return
    }
   
        fmt.Println(fi.Sys())
      
    
} 
