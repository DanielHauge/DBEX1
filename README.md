# Database excersise 1 (Key-Value Stores)
This is a mini project for the database course. The excersise is based off the description in the buttom of this [ressource](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/01-Intro_to_DB.ipynb)
## Description
This excersise goal is to build a simple database system in any language. I choose to build it in golang. The excersise needs to have implemented a hashmap-based index, functionality to store data on disk in binary file and read the data again to reinstantiate the database after a shutdown.

The result is a go library(SimpleGoDB.go) which can be used in any go program. Furthermore a simple CLI example program (Main.go)

------------------
## How to run it (example program)
- Pre-conditions: Have go installed. If you do not have go installed, install it [HERE](https://golang.org/dl/)
To run the example program, simply start by clonning the repo. From your prefered terminal. (I'm using git bash)
```
git clone https://github.com/Games-of-Threads/DBEX1-DFH.git
```
Then, go to the directory of the repository, and build and run with the following command.
```
go run Main.go SimpleGoDB.go
```
Now the program should be running. The example program is a simple command line interface to use the library. The commands is as follows:

- This command will initialize the database. If put in a database allready made, it will load the data, if not. It will create empty database, however this will not save it. (No point saving empty database. 

```initsdb [dbname]```



- This command will save the database to binaryfile on the default directory (usualy the same place the programfiles are). but will however not close down the program

```savesdb```

- This will save the data in the hashmap. the key must be integer. eg. put 9001 it's over nine thousand!!!

```put [key] [value]```

- This will print the key and value of the given value. eg. view 9001

```view [key]```

- This will list all data. Probably not so smart if the dataset becomes huge

```viewall```

- This command will delete the entry at the given key. eg. delete 9001

```delete [key]```

- This command will gracefully shut down the pgoram, but will save the database before doing so

```exit```

-------------
Testcase:

Try the following:
- Start the example program.
- Initialize a database (eg. Testdatabase)
- put in a few data points
- view all the data plotted in.
- exit by exit command
- Start the program up again and initialize same database
- view the previusly plotted data.
- see if you can delete an entry and save it with save command, and see if you can see the changes to the binary file with hexdump. (If you don't have hexdump for windows follow this [guide](https://superuser.com/questions/701141/how-to-add-more-commands-to-git-bash))

-----------------------
## How to integrate it (Use it in your go project)
To integrate this simple database into a go project. Simply add the SimpleGoDB.go to the project directory, and build and run it with the project.

The library comes with a struct to be instanced.
```go
type SDB struct {
	Name string
	DB map[int]string
}
```
This struct is basicly the hashmap and the filename. The name can be changed by the application(code) if demanded, this entails it will have a different filename when saved.

It also comes with a function to be called from the instanced SDB object.
```go
func (sdb *SDB) Save()
```
This will basicly just encode it with gob to binary and save it to disk.

Another function.
```go
func InitSDB(DBName string)SDB
```
This function call will instanziate a database. Load data if the filename provided allready exists, or instanziate an empty SDB object with only a filename which was provided.

It has 3 suplement functions not realy meant to be used by anyone else than the 2 previus functions. But made available if the need arose.

```go
func ReadFile(filename string) ([]byte, error)
```

```go
func WriteFile(filename string, data []byte) error
```

```go
func Exists(name string) bool
```

These functions, checks if file allready exists or not, write and read binary files with byte arrays.

example of using the library

```go
///// To initiaze DB
SimpleDB := InitSDB("MyFirstSimpleDB")
///// storing data
SimpleDB.DB[3] = "Some value as a string"
/////// Printing data that correspond to key 3
println(SimpleDB.DB[3])
/////// Fetching value from key and puts into variable
myvariable := SimpleDB.DB[9001]
////// Save the database to binaryfile
SimpleDB.Save()
```

## Why and when you should use this, and vice versa.
a few pro's and cons of this type of approach to a database.

Advantages | Disadvantages
----------------------- | ------------------------------------------
Fast lookup and insertion during runtime	| Slow during save to disk (Does not edit changed data, it deletes previus and saves new)
Simple to use   | Database only convinently accesible from program it is integrated with
Types can be changed easily | Needs to store data in memory

Here's a few critiria where it wouldn't be a bad idea to use this based off the pro's and cons.
- when you either don't have to much data or don't care about time when exiting/saving the data to disk.
- when the data doesn't need to be accesed by anyone other than the program.
- When you need/want fast lookup and insertion during runtime of a program.
- When you have no problems with memory usage
