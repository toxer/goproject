package main

import(
"fmt"
"time"
"strconv"
"os"

)

var maxSizeBeforeRotation int = 2000
var maxFileNumber int = 5
var logPath string ="/tmp/log"
var baseName string="server.log"
var millisecondBetweenWrite time.Duration = 100

func main() {
	//data
	var lineNumber int  = 0
	//pulisce la directory
	os.RemoveAll(logPath)
	fmt.Print("Rimosso "+logPath)
	
	os.Mkdir(logPath,0777)
	fmt.Println("Creato "+logPath)
		time.Sleep(10 * time.Millisecond)

	

	for{

		//eliminazione dell'ultimo file
		

		//cancello l'ultimo file
		os.Remove(logPath+"/"+baseName+"_"+strconv.Itoa(maxFileNumber))

		//inizio lo spostamento dei file
		for k := maxFileNumber; k >0;k--{
			oldSuffix :=""
			newSuffix:="_"+strconv.Itoa(k)
			if (k >1){
				oldSuffix="_"+strconv.Itoa((k-1))
			}

			os.Rename(logPath+"/"+baseName+oldSuffix,logPath+"/"+baseName+newSuffix)
			fmt.Println("Move "+logPath+"/"+baseName+oldSuffix+" to "+logPath+"/"+baseName+newSuffix)
		}

		//scrivo nel file di testa
		
		f, _ := os.Create(logPath+"/"+baseName)
		for i := 0; i < maxSizeBeforeRotation;{
			lineNumber++;
			t:= time.Now()
			s:=fmt.Sprint(t.Format(time.StampMicro))
			b:=[]byte(s+" Linea numero "+strconv.Itoa(lineNumber)+"\n")
			f.Write(b);
			i+=len(b)
			time.Sleep(millisecondBetweenWrite * time.Millisecond)

		}
		f.Write([]byte("--------File rotated-----------"));
		f.Sync()
		f.Close()
		

	}
	
}