package logstructure
import("os"
"fmt")



type Log struct{
	Inode,NextInode,PreviousInode uint64
	Index int
	CurrentName string	
	FilePointer os.FileInfo
}

type Logs []Log

func (slice Logs) Len() int {
	return len(slice)
}

func (slice Logs) Less(i, j int) bool {
	return slice[i].Index < slice[j].Index
}

func (slice Logs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (this Log) String() string {
	return fmt.Sprint(this.PreviousInode)+"-->"+this.CurrentName+"="+fmt.Sprint(this.Inode)+"-->"+fmt.Sprint(this.NextInode)+"\n"

}