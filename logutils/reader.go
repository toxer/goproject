package logutils
/*import("os"
"strings"
"io/ioutil"
"encoding/json"

)
func ReadDirectory(configurationFile string,dirPath string,partBeforeId string,partAfterId string)(string){
	//logs,mapInodeFile,mapNameInode := CreateLogList(dirPath,partBeforeId,partAfterId)
	logs,_,_ := CreateLogList(dirPath,partBeforeId,partAfterId)
	if (len(logs)==0){
		return
	}

	//cerco se esiste il file di configurazione
	confFile,err:=os.Open(configurationFile)
	if (err != nil){
		//creo il file
	    confFile,_=os.Create(configurationFile);
		confFile.Sync()
	
		
	}
	

	//effettuo la lettura del file di condfigurazione
	confData,_:=ioutil.ReadFile(confFile.Name())
	//se non è vuoto deincapsulo in json
	if(len(strings.TrimSpace(string(confData[:])))!=0){
		//deincapsulo il file nella struttura
	confStruct := new (ConfStruct)
	//controllo che l'inode del file sia anocra quello di testa


	err := json.Unmarshal(ioUtil.ReadFile(confFile.Name()),confStruct)
	if(err != nil)
	{
		panic
	}

		
	}else{

		//se vuoto, parto dall'inizio dell'ultimo file e salvo la struttura
		confStruct.Inode = logs[len(logs)-1].Inode
		confStruct.FileName = logs[len(logs)-1].CurrentName

	}





	//salvo la configurazione attuale
	b,err := json.Marshal(l)
	

	return logs[0].CurrentName
}*/