package filescanner
import(
//"fmt"
"io/ioutil"
"syscall"
"encoding/json"
"os"
)


type Log struct{
	Inode,NextInode,PreviusInode uint64
	index int
	CurrentName string	
}





// metodi per ordinare i file di log in base al nome

func (s Log) Len() int {
    return s.index
}

func (s Log) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s Log) Less(i, j int) bool {
	//trovo il numero finale del log
	 finalI = 0;
	 finalJ = 0;
	 
    return s[i].index < s[j].index
}

func (l Log )ToJson()(string,error){
	b,err := json.Marshal(l)
	return string(b[:]),err
}


func ToJson(l *Log)(string,error){
	b,err := json.Marshal(l)
	return string(b[:]),err
}
func ToStructure(jsonString string)(*Log,error){
	
	l := new (Log)
	err := json.Unmarshal([]byte(jsonString),l)
	return l,err
}


//recupera dal file di configurazione l'ultima struttura che indica lo stato della lettura
func CreateOrRetreiveConfigFile(filePath string)(*[]Log){
	return nil
}


//restituisce la mappa aggiornata e il nome del file in base all'inode passato


//compilo una lista di file da leggere

func FileNameByInode(dirPath string,inode uint64)(string){
	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		if (f.Sys().(*syscall.Stat_t).Ino==inode){
			return f.Name()
		}	
	}
	return ""
}


//resistuisce la mappa di inode-nomeFile
func InodeFileMapName(dirPath string)(map[uint64]string){
	inodeFileMap := make(map[uint64]string)

	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		inodeFileMap[(f.Sys().(*syscall.Stat_t).Ino)]=f.Name()
	}

	
	return inodeFileMap
}


//resistuisce la mappa di inode-puntatoreAFile
func InodeFileMapPointer(dirPath string)(map[uint64]*os.File){
	inodeFileMap := make(map[uint64]*File)

	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		inodeFileMap[(f.Sys().(*syscall.Stat_t).Ino)]=os.Open(f)
	}

	
	return inodeFileMap
}




