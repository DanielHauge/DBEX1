# Database excersise 1 (Key-Value Stores)
This is a mini project for the database course. The excersise is based off the description in the buttom of this [ressource](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/01-Intro_to_DB.ipynb)
## Description
This excersise goal is to build a simple database system in any language. I choose to build it in golang. The excersise needs to have implemented a hashmap-based index, functionality to store data on disk in binary file and read the data again to reinstantiate the database after a shutdown.

The result is a go library(SimpleGoDB.go) which can be used in any go program. Furthermore a simple CLI example program (Main.go)

------------------
## How to run it (example program)
- Pre-conditions: Have go installed. If you do not have go installed follow this Guide:LINK
To run the example program, simply start by clonning the repo. From your prefered terminal.
```
git clone https://github.com/Games-of-Threads/DBEX1-DFH.git
```
Then, go to the directory of the repository, and build and run with the following command.
```
go run Main.go SimpleGoDB.go
```
Now the program should be running. The example program is a simple command line interface to use the library. The commands is as follows:

```initsdb [dbname]```

- This command will initialize the database. If put in a database allready made, it will load the data, if not. It will create empty database, however this will not save it. (No point saving empty database. 

```savesdb```

- This command will save the database to binaryfile on the default directory (usualy the same place the programfiles are). but will however not close down the program
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

## How to integrate it (Use it in your go project)
Blabla
## Why and when you should use this, and vice versa.
Blabla
