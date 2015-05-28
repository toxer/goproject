package filemapgenerator

import (
	"fmt"
	"os"
)

type Log struct {
	Inode                uint64
	FileSize, ByteReaded int64
	Fp                   *os.File
}

func (this Log) String() string {
	return "Inode: " + fmt.Sprint(this.Inode) + " ByteReaded" + fmt.Sprint(this.ByteReaded) + " FileSize" + fmt.Sprint(this.FileSize)

}
