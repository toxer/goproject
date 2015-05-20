package main
import (
"fmt"
s "github.com/toxer/goproject/filescanner")


func main() {
	//l := new(s.Log)
	//fmt.Println(s.InodeFileMap("/tmp"))
	//name:=s.FileNameByInode("/tmp",10878992);

	l := new(s.Log)	
	l.Inode = 1
	l.NextInode=2
	l.PreviusInode=3
	l.CurrentName="test"
	json,_ := (l.ToJson())
	fmt.Println(json)
	fmt.Println(s.ToStructure(json))

	
	

} 
