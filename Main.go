package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

var DB SDB

func main() {

	CLI := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter Command. \nAvailable Commands: \n initsdb [sdb name]      savesdb      put [key] [value...]       view [key]       viewall      delete [key]        exit\n")
	for {

		input, _ := CLI.ReadString('\n')
		command := strings.Fields(input)

		switch command[0] {
		case "initsdb":
			DB = InitSDB(command[1])

		case "savesdb":
			DB.Save()

		case "put":
			key,_ := strconv.Atoi(command[1])
			value := strings.Join(command[1:]," ")
			DB.DB[key] = value

		case "view":
			key,_ := strconv.Atoi(command[1])
			fmt.Println(DB.DB[key])

		case "viewall":
			for _, value := range DB.DB{
				println(value)
			}


		case "delete":
			key,_ := strconv.Atoi(command[1])
			delete(DB.DB, key)

		case "exit":
			DB.Save()
			os.Exit(2)

		default:
			println("Did not recognize the command.")
			fmt.Print("\nEnter Command. \nAvailable Commands: \n initsdb [sdb name]      savesdb      put [key] [value...]       view [key]       viewall      delete [key]        exit\n")
		}




	}

	/*
	///// To initiaze DB
	DB := InitSDB("MyFirstSimpleDB")

	///// storing data
	DB.DB[3] = "Testing Testing ->"

	/////// Printing data
	println(DB.DB[3])

	/////// Fetching value from key
	myvariable := DB.DB[9001]

	////// Save the database to binaryfile
	DB.Save()
	*/
}
