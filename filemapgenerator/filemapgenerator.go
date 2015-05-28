package filemapgenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func GetLogFiles(dirPath string, configurationFilePath string) []Log {

	//scansiono la directory
	files, err := ioutil.ReadDir(dirPath)
	//estraggo il file di configurazione
	a, err := os.Open(configurationFilePath)

	strSaved := new([]Log)

	//lo creo se non esiste
	if err != nil {
		a.Close()
		fmt.Println("Creazione del file vuoto " + configurationFilePath)
		confFile, _ := os.Create(configurationFilePath)
		confFile.Sync()
		confFile.Close()

	}
	str, _ := ioutil.ReadFile(configurationFilePath)
	fmt.Println(configurationFilePath)
	if len(str) != 0 {
		fmt.Println("Caricamento struttura preesistente")
		strSaved = new([]Log)
		//carico la struttura presente
		er := json.Unmarshal(str, &strSaved)
		if er != nil {
			fmt.Print(er)
		}

	} else {
		fmt.Println("Nessuna struttura preesistente")
	}

	//estraggo i logs
	logs := make([]Log, len(files))

	if len(files) == 0 {
		return make([]Log, 0)
	}

	//riempio le strutture
	index := 0
	for _, f := range files {
		l := new(Log)

		l.Fp, err = os.Open(dirPath + "/" + f.Name())
		if err != nil {
			panic(fmt.Sprint(err))
		}
		l.Inode = f.Sys().(*syscall.Stat_t).Ino
		//se la struttra salvata non è nulla
		//cerco in base all'inode del file
		//e ricavo i byte letti
		l.ByteReaded = 0
		if strSaved != nil {
			for _, lSav := range *strSaved {
				if lSav.Inode == l.Inode {
					l.ByteReaded = lSav.ByteReaded
					break
				}
			}
		}

		l.FileSize = f.Size()
		logs[index] = *l
		index++

	}
	return logs

}

func SaveStructure(configurationFilePath string, logs []Log, closeAllFile ...bool) {
	//persisto la struttura trasformandola in json e scrivendo il file
	//che deve esistere

	b, err := json.Marshal(logs)
	if err == nil {
		fmt.Println("Salvataggio struttura in corso: " + fmt.Sprint(logs))
		ioutil.WriteFile(configurationFilePath, b, 0777)
	} else {
		panic(fmt.Sprint(err))
	}
	for _, l := range logs {
		if l.Fp != nil && ((len(closeAllFile) > 0 && closeAllFile[0] == true) || len(closeAllFile) == 0) {
			fmt.Print("Close " + fmt.Sprint(l.Inode))
			l.Fp.Close()
		}
	}

}

func (l Log) Save(configurationFilePath string) {
	//do per scontato che esista il file di configurazione
	fconf, err := os.Open(configurationFilePath)
	if err != nil {
		fmt.Println(err)
		fconf.Close()
		return
	}

	//carico la struttura salvata
	str, _ := ioutil.ReadFile(configurationFilePath)
	strSaved := make([]Log, 1)
	if len(str) != 0 {

		//carico la struttura presente
		er := json.Unmarshal(str, &strSaved)
		if er != nil {
			fmt.Print(er)
		}
		fmt.Println("SAVE: Caricamento struttura preesistente " + fmt.Sprint(strSaved))

	} else {
		fmt.Println("Save creo struttura e salvo")
		//creo la struttra e salvo
		os.Create(configurationFilePath)
		logs := make([]Log, 1)
		logs[0] = l
		SaveStructure(configurationFilePath, logs, false)
		fconf.Close()
		return
	}
	//qui carico e aggiorno la struttura preesistente

	for _, ls := range strSaved {
		if ls.Inode == l.Inode {
			ls.ByteReaded = l.ByteReaded
			ls.FileSize = l.FileSize
			ls.Fp = l.Fp

			fmt.Println("Aggiorno struttura e salvo " + fmt.Sprint(strSaved))
			//creo la struttra e salvo
			//salvo la struttura
			SaveStructure(configurationFilePath, strSaved, false)
			fconf.Close()
			return
		}
	}

	//qui il nodo non era presente nella struttura, lo devo aggiungere
	fmt.Println("Aggiungo a struttura e salvo")
	//creo la struttra e salvo
	newStr := append(strSaved, l)
	fmt.Println("Nuova struttura: " + fmt.Sprint(newStr))
	//aggiungo e salvo

	SaveStructure(configurationFilePath, newStr, false)

	fconf.Close()
}
