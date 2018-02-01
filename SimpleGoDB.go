package main

import (
	"os"
	"bytes"
	"log"
	"bufio"
	"encoding/gob"
)

type SDB struct {
	Name string
	DB map[int]string
}

func (sdb *SDB) Save() {
	log.Println("SDB: Saving Data")
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(sdb.DB)
	if err != nil{
		log.Fatal(err)
	}
	err = WriteFile(sdb.Name, buf.Bytes())
	if err != nil{
		log.Fatal(err)
	}
}



func InitSDB(DBName string)SDB{
	p := SDB{DBName, map[int]string{}}
	p.Name = DBName
	if Exists(DBName){
		log.Println("SDB: This database allready exists, loading data")

		data, err := ReadFile(DBName)
		if err != nil{
			log.Fatal(err)
		}


		dec := gob.NewDecoder(bytes.NewReader(data))
		gob.Register(map[int]string{})
		dec.Decode(&p.DB)


	}else {
		log.Println("SDB: This database did not exist, database initizalised.")
		p.DB = map[int]string{}
	}
	return p
}

func ReadFile(filename string) ([]byte, error){
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_,err = bufr.Read(bytes)

	return bytes, err
}

func WriteFile(filename string, data []byte) error{

	if Exists(filename){
		err := os.Remove(filename)
		if err != nil{
			log.Fatal(err)
		}
	}
	file, err := os.Create(filename)
	defer file.Close()


	bufw := bufio.NewWriter(file)
	_, err = bufw.Write(data)
	err = bufw.Flush()

	return err
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}