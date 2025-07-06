package code

import (
	"bufio"
	"bytes"
	"code_paro/models"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const maxThreads = 4;
var wg sync.WaitGroup


func readSection(filename string, start int64, end int64, ch chan<-models.DataSet, sectionId int){
	defer wg.Done()

	file, err:= os.Open(filename)
	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	file.Seek(start,0)
	reader:= bufio.NewReader(file)
	var buffer bytes.Buffer
	var bytesRead int64 = 0

	for{
		byte, err := reader.ReadByte()
		if err!=nil{
			if err == io.EOF{
				break
			}
			fmt.Println("Read error",err)
			return
		}

		bytesRead++
		buffer.WriteByte(byte)

		if start+bytesRead > end && (byte == ' ' || byte == '\n') {
			break
		}

	}

	byteArr:= strings.Fields(buffer.String())
	localArr:= make(map[int]bool)

	for _, b:= range byteArr{
		if num, err:= strconv.Atoi(b); err == nil{
			localArr[num] = true;
		}
	}

	ch <- models.DataSet{Data: localArr, SectionID: sectionId}
	
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
	sections:= fileSize/ int64(maxThreads)

	ch:= make(chan models.DataSet, maxThreads)

	for i:=0; i<maxThreads; i++{
		start:= int64(i) * sections
		end:= start + sections
		wg.Add(1)
		go readSection(filename, start, end, ch, i)

	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	results:= make([]models.DataSet, maxThreads)
	for data := range ch{
		results[data.SectionID] = data;
	}
}

func Q1(){
	

}