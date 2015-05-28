package logutils

import (
	"io/ioutil"
	"os"
	//"fmt"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

//questo medodo scansiona una directory
//e compila un array di Log in base al path
//tipo path: parte_iniziale{ID}parte_finale
//il campo id viene usato per trovare l'indice del file
//se non viene trovato l'id vale 0

//viene anche restituita una mappa di puntatori ai file APERTI!!!

func CreateLogList(dirPath string, partBeforeId string, partAfterId string) ([]Log, map[uint64]*os.File, map[string]uint64) {
	//per prima cosa scansiono la directory e costruisco la mappa dei file
	//completando i soli id
	inodeFileMap := make(map[uint64]*os.File)
	stringInodeMap := make(map[string]uint64)
	files, _ := ioutil.ReadDir(dirPath)
	logs := make(Logs, len(files))
	i := 0
	for _, f := range files {

		inode := f.Sys().(*syscall.Stat_t).Ino
		l := new(Log)
		l.Inode = inode
		l.CurrentName = f.Name()
		tmpId := strings.Replace((strings.Replace(f.Name(), partBeforeId, "", 1)), partAfterId, "", 1)
		if tmpId == f.Name() {
			l.Index = 0
		} else {
			l.Index, _ = strconv.Atoi(tmpId)
		}
		logs[i] = *l
		fp, _ := os.Open(f.Name())
		inodeFileMap[inode] = fp
		stringInodeMap[f.Name()] = inode
		i++

	}

	//effettuo un ordinamento per indice
	sort.Sort(logs)

	//in base all'ordinamento e alla mappa completo la linked list
	//nell'array dei logs

	for i := 0; i < len(logs); i++ {
		if i > 0 {
			logs[i].PreviousInode = logs[i-1].Inode
		}
		if i < len(logs)-1 {
			logs[i].NextInode = logs[i+1].Inode
		}
	}

	return logs, inodeFileMap, stringInodeMap
}
