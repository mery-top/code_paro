package code

import(
	"fmt"
	"os"
	"sync"
	"log"
	"code_paro/models"
)

const maxTheards = 4;
var wg sync.WaitGroup


func readSection(filename string, start int64, end int64, ch chan<-models.DataSet){
	

}


func callThreads(){
	filename:= "number.txt"
	file, err:= os.Open(filename)
	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	fileStat, err:= file.Stat()
	if err!=nil{
		log.Fatal(err)
	}

	//divide sections
	fileSize:= fileStat.Size()
	sections:= fileSize/ int64(maxTheards)

	ch:= make(chan models.DataSet, maxTheards)

	for i:=0; i<maxTheards; i++{
		start:= int64(i) * sections
		end:= start + sections
		wg.Add(1)
		go readSection(filename, start, end, ch)

	}

	go func(){
		wg.Wait()
		close(ch)
	}()
}

func Q1(){
	

}