# CRAM Cli 

## What is dbsock?

In order to increase efficiency we create a program that creates a unix doamin socket that maintains the database conection even while the cli isnt running. Dbsock is the program that creates that socket and connects it to the database.

## Building dbsock

In order to build dbsock, from the root of the project enter these commands:

```shell
$ cd dbsock
$ go build -o dbsock ./dbsock.go
```

## Building the main program

From the root of the project do

```shell
$ go build -o main main.go
```
