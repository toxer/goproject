package main

import (
	"fmt"
	ls "github.com/toxer/goproject/filemapgenerator"
)

func main() {

	fp := "/tmp/conf/save.conf"
	l := ls.GetLogFiles("/tmp/log", fp)
	l[0].FileSize = 7
	//ls.SaveStructure(fp, l)

	//l := ls.ReadDirectory("/tmp/conf/logsave.cfg", "/tmp/log", "server.log_", "")
	ll := new(ls.Log)
	ll.Inode = 1234
	ll.FileSize = 90000
	ll.ByteReaded = 2
	ll.Save(fp)
	fmt.Println(ls.GetLogFiles("/tmp/log", fp))

}
